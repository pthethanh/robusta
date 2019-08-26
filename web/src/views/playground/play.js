export const playground = {
  blocks: [
    {
      'type': 'header',
      'data': {
        'text': 'Why Telegram is the best messenger',
        'level': 2
      }
    },
    {
      'type': 'checklist',
      'data': {
        'items': [
          {
            'text': 'This is a block-styled editor',
            'checked': true
          },
          {
            'text': 'Clean output data',
            'checked': false
          },
          {
            'text': 'Simple and powerful API',
            'checked': true
          }
        ]
      }
    },
    {
      'type': 'delimiter',
      'data': {}
    },
    {
      'type': 'embed',
      'data': {
        'service': 'youtube',
        'source': 'https://www.youtube.com/watch?v=qa0nCMVk2FE',
        'embed': 'https://www.youtube.com/embed/qa0nCMVk2FE',
        'width': 580,
        'height': 320,
        'caption': 'My Life'
      }
    },
    {
      'type': 'simpleImage',
      'data': {
        'url': 'https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg',
        'caption': 'Roadster // tesla.com',
        'withBorder': true,
        'withBackground': true,
        'stretched': true
      }
    },
    {
      'type': 'linkTool',
      'data': {
        'link': 'https://codex.so',
        'meta': {
          'title': 'CodeX Team',
          'site_name': 'CodeX',
          'description': 'Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.',
          'image': {
            'url': 'https://codex.so/public/app/img/meta_img.png'
          }
        }
      }
    },
    {
      'type': 'raw',
      'data': {
        'html': '<div style="background: #000; color: #fff; font-size: 30px; padding: 50px;">Any HTML code</div>'
      }
    },
    {
      'type': 'paragraph',
      'data': {
        'text': 'Check out our projects on a <a href="https://github.com/codex-team">GitHub page</a>.'
      }
    },
    {
      'type': 'warning',
      'data': {
        'title': 'Note:',
        'message': 'Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.'
      }
    },
    {
      'type': 'quote',
      'data': {
        'text': 'The unexamined life is not worth living.',
        'caption': 'Socrates',
        'alignment': 'left'
      }
    },
    {
      'type': 'table',
      'data': {
        'content': [ ['Kine', '1 pcs', '100$'], ['Pigs', '3 pcs', '200$'], ['Chickens', '12 pcs', '150$'] ]
      }
    }
  ]
}
