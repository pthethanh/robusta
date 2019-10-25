<template>
  <div class="resources">
    <el-row type="flex" justify="center">
      <el-col :xs="24" :sm="24" :md="18" :lg="14" :xl="14">
        <h2 class="section-title">{{ $t('resource.articles') }}</h2>
        <div v-for="(article, index) in articles" :key="index">
          <h3>{{ $t('resource.' + article.level) }}</h3>
          <ul class="block">
            <li v-for="(category, idx) in article.categories" :key="idx">
              <span>{{category.name}}:&nbsp;</span>
              <ul>
                <li v-for="(link, idx) in category.links" :key="idx">
                  <span>{{link.name}}:&nbsp;</span>
                  <el-link :href="link.link" target="_blank" type="primary">{{link.link}}</el-link>
                </li>
              </ul>
            </li>
            <li v-for="(link, idx) in article.links" :key="idx">
              <span>{{link.name}}:&nbsp;</span>
              <el-link :href="link.link" target="_blank" type="primary">{{link.link}}</el-link>
            </li>
          </ul>
        </div>
        <h2 class="section-title">{{ $t('resource.books') }}</h2>
        <ul class="block">
          <li v-for="(book,index) in books" :key="index">
            <el-link :href="book.link" target="_blank" type="primary">{{book.name}}</el-link>
          </li>
        </ul>
      </el-col>
    </el-row>
  </div>
</template>

<script>
/* eslint-disable */
export default {
  data () {
    return {
      activeNames: ['1'],
      books: [{
          name: 'The Go Programming Language - Donovan, Alan A. A',
          link: 'https://drive.google.com/open?id=1-PNY3PXfwKnaOdFHeOx0JuC8NnCVQmgF',
          image: 'https://images-na.ssl-images-amazon.com/images/I/516R4ZoMqBL._SX402_BO1,204,203,200_.jpg',
        },
        {
          name: 'Go in Practice',
          link: 'https://drive.google.com/open?id=1QhsqRZOsO9zd3aEiq_OE8nKDNEqxVV40',
        },
        {
          name: 'Build web application with Go',
          link: 'https://github.com/astaxie/build-web-application-with-golang',
          image: 'https://golang.org/lib/godoc/images/footer-gopher.jpg'
        },
        {
          name: 'Go Advanced Programming',
          link: 'https://drive.google.com/open?id=1lW8EniL8LkLtbCxRIvc_x3ELZCjPcYho',
          image: 'https://golang.org/lib/godoc/images/footer-gopher.jpg'
        },
        {
          name: 'Go programming language secure coding practices guide (FREE)',
          link: 'https://checkmarx.gitbooks.io/go-scp/',
          image: 'https://golang.org/lib/godoc/images/footer-gopher.jpg'
        }
      ],
      articles: [{
          level: 'beginner',
          links: [{
              name: 'The Tour of Go',
              link: 'https://tour.golang.org/welcome/1'
            },
            {
              name: 'Go by example',
              link: 'https://gobyexample.com/'
            },
            {
              name: 'Install Go',
              link: 'https://golang.org/doc/install'
            },
            {
              name: 'How to write Go code',
              link: 'https://golang.org/doc/code.html'
            },
            {
              name: 'Effective Go',
              link: 'https://golang.org/doc/effective_go.html'
            },
            {
              name: 'Go FAQ',
              link: 'https://golang.org/doc/faq'
            },
            {
              name: 'Go Talks',
              link: 'https://talks.golang.org'
            },
            {
              name: 'Awesome Libraries',
              link: 'https://github.com/avelino/awesome-go'
            }
          ]
        },
        {
          level: 'intermediate',
          categories: [{
              name: 'Array, Slice',
              links: [{
                  name: 'Go slices usage and internals',
                  link: 'https://blog.golang.org/go-slices-usage-and-internals'
                },
                {
                  name: 'Slices',
                  link: 'https://blog.golang.org/slices'
                }
              ]
            },
            {
              name: 'Strings',
              links: [{
                name: 'Strings',
                link: 'https://blog.golang.org/strings'
              }]
            },
            {
              name: 'Errors',
              links: [{
                  name: 'Errors are values',
                  link: 'https://blog.golang.org/errors-are-values'
                },
                {
                  name: 'Error handling in Go',
                  link: 'https://blog.golang.org/error-handling-and-go'
                }
              ]
            },
            {
              name: 'Interface',
              links: [{
                name: 'Gift decoder exercise in Go interface',
                link: 'https://blog.golang.org/gif-decoder-exercise-in-go-interfaces'
              }]
            },
            {
              name: 'Goroutines',
              links: [{
                  name: 'Sharing by communicating',
                  link: 'https://blog.golang.org/share-memory-by-communicating'
                },
                {
                  name: 'Concurrency is not parallelism',
                  link: 'https://blog.golang.org/concurrency-is-not-parallelism'
                },
                {
                  name: 'Advanced Go concurrency patterns',
                  link: 'https://blog.golang.org/advanced-go-concurrency-patterns'
                },
                {
                  name: 'Go Concurrency Patterns: Context',
                  link: 'https://blog.golang.org/context'
                },
                {
                  name: 'Go concurrency patterns timing out, moving on',
                  link: 'https://blog.golang.org/go-concurrency-patterns-timing-out-and'
                }
              ]
            },
            {
              name: 'HTTP',
              links: [{
                name: 'HTTP package in Go doc',
                link: 'https://godoc.org/net/http'
              }]
            },
            {
              name: 'Testing',
              links: [{
                  name: 'Testing package',
                  link: 'https://godoc.org/testing'
                },
                {
                  name: 'HTTP Testing',
                  link: 'https://godoc.org/net/http/httptest'
                },
                {
                  name: 'Table driven testing',
                  link: 'https://godoc.org/testing'
                },
                {
                  name: 'Using Subtests and Sub - benchmarks',
                  link: 'https://blog.golang.org/subtests'
                }
              ]
            },
            {
              name: 'Database',
              links: [{
                  name: 'Golang database tutorial',
                  link: 'http://go-database-sql.org'
                },
                {
                  name: 'SQLX',
                  link: 'http://jmoiron.github.io/sqlx/'
                }
              ]
            }
          ]
        },
        {
          level: 'advanced',
          links: [{
              name: 'Go Proverbs',
              link: 'https://go-proverbs.github.io'
            },
            {
              name: 'Laws of reflection',
              link: 'https://blog.golang.org/laws-of-reflection'
            },
            {
              name: 'Interface',
              link: 'https://research.swtch.com/interfaces'
            },
            {
              name: 'Profiling Go programs',
              link: 'https://blog.golang.org/profiling-go-programs'
            },
            {
              name: "Getting to Go: The Journey of Go's Garbage Collector",
              link: 'https://blog.golang.org/ismmkeynote'
            },
            {
              name: "A Quick Guide to Go's Assembler",
              link: 'https://golang.org/doc/asm'
            },
          ]
        }
      ]
    }
  }
}
</script>

<style lang="scss" scoped>
@import '@/styles/variables.scss';

.resources {
  font-family: 'Open Sans', sans-serif;
  font-size: 0.95em;
  padding: 20px 20px;
  background-color: $backgroundColorSecondary;

  .block {
    display: block;
    box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
    background-color: $backgroundColorInfo;
    border-radius: 5px;
  }
}
</style>
