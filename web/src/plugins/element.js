import Vue from 'vue'
import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'

import {
  Dialog,
  Dropdown,
  DropdownMenu,
  DropdownItem,
  Menu,
  Submenu,
  MenuItem,
  MenuItemGroup,
  Input,
  Button,
  ButtonGroup,
  Popover,
  Form,
  FormItem,
  Tag,
  Alert,
  Icon,
  Row,
  Col,
  Progress,
  Spinner,
  Card,
  Steps,
  Step,
  Container,
  Header,
  Aside,
  Main,
  Footer,
  Timeline,
  TimelineItem,
  Link,
  Divider,
  Image,
  PageHeader,
  Loading,
  MessageBox,
  Message,
  Notification,
  Avatar,
  InfiniteScroll,
  Tabs,
  TabPane,
  Table,
  TableColumn,
  Select,
  Option,
  Drawer,
  Badge
  // Breadcrumb,
  // BreadcrumbItem,
  // DatePicker,
  // TimeSelect,
  // TimePicker,
  // InputNumber,
  // Radio,
  // RadioGroup,
  // RadioButton,
  // Checkbox,
  // CheckboxButton,
  // CheckboxGroup,
  // Switch,
  // OptionGroup,
  // Pagination,
  // Autocomplete,
  // CascaderPanel,
  // Calendar,
  // Backtop,
  // Carousel,
  // CarouselItem,
  // Collapse,
  // CollapseItem,
  // Cascader,
  // ColorPicker,
  // Transfer,
  // Rate,
  // Upload,
  // Slider,
  // Tree,
  // Tooltip,
} from 'element-ui'

locale.use(lang)

Vue.use(Dialog)
Vue.use(Dropdown)
Vue.use(DropdownMenu)
Vue.use(DropdownItem)
Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItem)
Vue.use(MenuItemGroup)
Vue.use(Input)
Vue.use(Button)
Vue.use(ButtonGroup)
Vue.use(Popover)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Tag)
Vue.use(Alert)
Vue.use(Icon)
Vue.use(Row)
Vue.use(Col)
Vue.use(Progress)
Vue.use(Spinner)
Vue.use(Card)
Vue.use(Steps)
Vue.use(Step)
Vue.use(Container)
Vue.use(Header)
Vue.use(Aside)
Vue.use(Main)
Vue.use(Footer)
Vue.use(Timeline)
Vue.use(TimelineItem)
Vue.use(Link)
Vue.use(Divider)
Vue.use(Image)
Vue.use(PageHeader)
Vue.use(Avatar)
Vue.use(InfiniteScroll)
Vue.use(Tabs)
Vue.use(TabPane)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Select)
Vue.use(Option)
Vue.use(Drawer)
Vue.use(Badge)
// Vue.use(Breadcrumb)
// Vue.use(BreadcrumbItem)
// Vue.use(Calendar)
// Vue.use(Backtop)
// Vue.use(CascaderPanel)
// Vue.use(Rate)
// Vue.use(Upload)
// Vue.use(Slider)
// Vue.use(Tree)
// Vue.use(Tooltip)
// Vue.use(Autocomplete)
// Vue.use(Pagination)
// Vue.use(InputNumber)
// Vue.use(Radio)
// Vue.use(RadioGroup)
// Vue.use(RadioButton)
// Vue.use(Checkbox)
// Vue.use(CheckboxButton)
// Vue.use(CheckboxGroup)
// Vue.use(Switch)
// Vue.use(OptionGroup)
// Vue.use(DatePicker)
// Vue.use(TimeSelect)
// Vue.use(TimePicker)
// Vue.use(Carousel)
// Vue.use(CarouselItem)
// Vue.use(Collapse)
// Vue.use(CollapseItem)
// Vue.use(Cascader)
// Vue.use(ColorPicker)
// Vue.use(Transfer)

Vue.use(Loading.directive)

Vue.prototype.$loading = Loading.service
Vue.prototype.$msgbox = MessageBox
Vue.prototype.$alert = MessageBox.alert
Vue.prototype.$confirm = MessageBox.confirm
Vue.prototype.$prompt = MessageBox.prompt
Vue.prototype.$notify = Notification
Vue.prototype.$message = Message
