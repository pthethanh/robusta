<template>
  <div class="home">
    <Article v-for="article in articles" :key="article.id" :article="article"></Article>
  </div>
</template>

<script>
import Article from '@/components/Article'
export default {
  components: {
    Article
  },
  data () {
    return {
      articles: [
        {
          id: '1',
          title: 'Hello Go',
          abstract: '<p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p>',
          content: '<p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p><p>Test</p>',
          created_by_id: 'thanhpham',
          created_by_name: 'thanh pham the',
          created_at: '',
          likes: 20,
          comments: 300,
          views: 560
        },
        {
          id: '2',
          title: 'Hello Go',
          abstract: 'Hello go from go team',
          content: 'Hello Go content',
          created_by_id: 'thanhpham',
          created_by_name: 'thanh pham the',
          created_at: '',
          likes: 20,
          comments: 300,
          views: 560
        },
        {
          id: '3',
          title: 'Hello Go',
          abstract: 'Hello go from go team',
          content: 'Hello Go content',
          created_by_id: 'thanhpham',
          created_by_name: 'thanh pham the',
          created_at: '',
          likes: 20,
          comments: 300,
          views: 560
        },
        {
          id: '4',
          title: 'Hello Go',
          abstract: 'Hello go from go team',
          content: 'Hello Go content',
          created_by_id: 'thanhpham',
          created_by_name: 'thanh pham the',
          created_at: '',
          likes: 20,
          comments: 300,
          views: 560
        }
      ]
    }
  }
}
</script>

<style>
  .home {
      margin: 0px 75px 0 75px;;
  }
</style>
