package policy

import (
	"context"
	"strings"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/db/mongodb"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/validator"

	"github.com/casbin/casbin"
	mongodbadapter "github.com/casbin/mongodb-adapter"
	"github.com/pkg/errors"
)

type (
	CasbinConfig struct {
		MongoDB    mongodb.Config
		ConfigPath string `envconfig:"CONFIG_PATH" default:"configs/casbin.conf"`
	}
	Service struct {
		enforcer *casbin.Enforcer
	}
)

// New return a new instance of policy service
func New(enforcer *casbin.Enforcer) (*Service, error) {
	enforcer.EnableAutoSave(true) // auto save...
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	return &Service{
		enforcer: enforcer,
	}, nil
}

// NewMongoDBCasbinEnforcer return new instance of cashout.Enforcer
// which use MongoDB as adapter.
func NewMongoDBCasbinEnforcer(conf CasbinConfig) *casbin.Enforcer {
	dialInfo := conf.MongoDB.DialInfo()
	adapter := mongodbadapter.NewAdapterWithDialInfo(dialInfo)
	enforcer := casbin.NewEnforcer(conf.ConfigPath, adapter)
	// TODO add watchers in the future to sync the policies between multiple nodes
	// see https://casbin.org/docs/en/watchers
	return enforcer
}

// AddPolicy add new policy
// For specific user:
// - s.AddPolicy("alice", "article_1", "read", allow)
// For group:
// - s.AddPolicy("group_admin", "article_1", "*", deny)
// - s.AddPolicy("group_admin", "article_2", "read", allow)
// - s.AddPolicy("group_admin", "article_3", "write", allow)
func (s *Service) AddPolicy(ctx context.Context, p types.Policy) error {
	_, err := s.enforcer.AddPolicySafe(p.Subject, p.Object, p.Action, p.Effect)
	return err
}

// AddGroupingPolicy add grouping policy...
// Example adding user alice into the group_admin
// that would make alice inherits all permissions from the group
// - e.AddGroupingPolicy("alice", "group_admin")
func (s *Service) AddGroupingPolicy(ctx context.Context, sub string, group string) error {
	_, err := s.enforcer.AddGroupingPolicySafe(sub, group)
	return err
}

// IsAllowed check if the sub is allowed to do the act on the obj
func (s *Service) IsAllowed(ctx context.Context, sub string, obj string, act string) bool {
	ok, err := s.enforcer.EnforceSafe(sub, obj, act)
	return err == nil && ok
}

// MakeOwner make the sub to be owner of the obj
func (s *Service) MakeOwner(ctx context.Context, sub string, obj string) error {
	return s.AddPolicy(ctx, types.Policy{
		Subject: sub,
		Object:  obj,
		Action:  types.PolicyActionAny,
		Effect:  types.PolicyEffectAllow,
	})
}

// AssignPolicy assign policy to a subject. Subject can be  user or role
func (s *Service) AssignPolicy(ctx context.Context, req AssignPolicyRequest) error {
	if err := validator.Validate(req); err != nil {
		return err
	}
	if err := s.validatePermission(ctx, req.Object, ActionPolicyUpdate); err != nil {
		return err
	}
	if err := s.AddPolicy(ctx, types.Policy{
		Subject: req.Subject,
		Object:  req.Object,
		Action:  req.Action,
		Effect:  req.Effect,
	}); err != nil {
		log.WithContext(ctx).Errorf("failed to add policy, err: %v", err)
		return errors.Wrap(err, "failed to add policy")
	}
	if req.Effect == types.PolicyEffectDeny {
		return nil
	}
	// cleanup existing old deny effect
	if _, err := s.enforcer.RemovePolicySafe(req.Subject, req.Object, req.Action, types.PolicyEffectDeny); err != nil {
		log.WithContext(ctx).Errorf("failed cleanup existing deny policy, err: %v", err)
		return errors.Wrap(err, "failed cleanup existing deny policy")
	}
	return nil
}

// AssignGroupPolicy assign user to roles
func (s *Service) AssignGroupPolicy(ctx context.Context, req AssignGroupPolicyRequest) error {
	if err := validator.Validate(req); err != nil {
		return err
	}
	if err := s.validatePermission(ctx, req.Group, ActionPolicyUpdate); err != nil {
		return err
	}
	if err := s.AddGroupingPolicy(ctx, req.Subject, req.Group); err != nil {
		log.WithContext(ctx).Errorf("failed to add group policy, err: %v", err)
		return errors.Wrap(err, "failed to add group policy")
	}
	return nil
}

// GetRoles get all available roles
func (s *Service) GetRoles(ctx context.Context) ([]string, error) {
	if err := s.validatePermission(ctx, Object, ActionPolicyUpdate); err != nil {
		return nil, err
	}
	roles := s.enforcer.GetAllRoles()
	return roles, nil
}

// GetUsersForRole return users of the given role
func (s *Service) GetUsersForRole(ctx context.Context, role string) ([]string, error) {
	if err := s.validatePermission(ctx, Object, ActionPolicyUpdate); err != nil {
		return nil, err
	}
	users := s.enforcer.GetRolesForUser(role)
	return users, nil
}

// FindPolicies return all policies match the given filters
func (s *Service) FindPolicies(ctx context.Context, req FindPolicyRequest) ([]types.Policy, error) {
	if err := s.validatePermission(ctx, Object, ActionPolicyUpdate); err != nil {
		return nil, err
	}
	policies := s.enforcer.GetPolicy()
	rs := make([]types.Policy, 0)
	for _, p := range policies {
		plc := types.Policy{
			Subject: p[0],
			Object:  p[1],
			Action:  p[2],
			Effect:  p[3],
		}
		matched := len(req.Subjects) == 0 // matched by default if not filter
		for _, sub := range req.Subjects {
			if plc.Subject == sub {
				matched = true
			}
		}
		if !matched {
			continue
		}
		matched = len(req.Objects) == 0 // matched by default if not filter
		for _, obj := range req.Objects {
			if plc.Object == obj {
				matched = true
			}
		}
		if !matched {
			continue
		}
		matched = len(req.Actions) == 0 // matched by default if not filter
		for _, act := range req.Actions {
			if strings.HasPrefix(plc.Action, act) {
				matched = true
				break
			}
		}
		if !matched {
			continue
		}
		rs = append(rs, plc)
	}
	return rs, nil
}

func (s *Service) RemovePolicy(ctx context.Context, req types.Policy) error {
	if err := validator.Validate(req); err != nil {
		return err
	}
	if err := s.validatePermission(ctx, req.Object, ActionPolicyUpdate); err != nil {
		return err
	}
	if _, err := s.enforcer.RemovePolicySafe(req.Subject, req.Object, req.Action, req.Effect); err != nil {
		log.WithContext(ctx).Errorf("failed to remove policy, err: %v", err)
		return errors.Wrap(err, "failed to remove policy")
	}
	return nil
}

func (s *Service) validatePermission(ctx context.Context, obj string, act string) error {
	user := auth.FromContext(ctx)
	if user == nil {
		return types.ErrUnauthorized
	}
	if !s.IsAllowed(ctx, user.UserID, obj, act) {
		return types.ErrUnauthorized
	}
	return nil
}

// ListActions return  all supported actions
func (s *Service) ListActions(ctx context.Context) ([]string, error) {
	user := auth.FromContext(ctx)
	if user == nil {
		return nil, types.ErrUnauthorized
	}
	if !s.IsAllowed(ctx, user.UserID, Object, ActionPolicyUpdate) {
		return nil, types.ErrUnauthorized
	}
	return []string{
		types.PolicyActionArticleCreate,
		types.PolicyActionArticleDelete,
		types.PolicyActionArticleUpdate,
		types.PolicyActionArticleRead,

		types.PolicyActionChallengeCreate,
		types.PolicyActionChallengeDelete,
		types.PolicyActionChallengeUpdate,
		types.PolicyActionChallengeRead,

		types.PolicyActionCommentCreate,
		types.PolicyActionCommentDelete,
		types.PolicyActionCommentUpdate,
		types.PolicyActionCommentRead,

		types.PolicyActionFolderCreate,
		types.PolicyActionFolderDelete,
		types.PolicyActionFolderUpdate,
		types.PolicyActionFolderRead,

		types.PolicyActionSolutionRead,
		types.PolicyActionSolutionReadListDetail,

		types.PolicyActionUserReadList,
	}, nil
}
