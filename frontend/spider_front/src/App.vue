<!-- <script setup>
import HelloWorld from './components/HelloWorld.vue'
import TheWelcome from './components/TheWelcome.vue'
</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />
    </div>
  </header>

  <main>
    <TheWelcome />
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style> -->


<!-- <template>  
  <div>  
    <h1>Movie Info</h1>  
    <table>
    <thead>
      <tr>
          <th>导演</th>  
          <th>编剧</th>  
          <th>演员</th>  
          <th>类型</th>  
          <th>地区</th>  
          <th>语言</th>  
          <th>上映日期</th>  
          <th>时长</th>  
          <th>别名</th>  
        </tr>  
      </thead>  
      <tbody>  
        <tr v-for="(movie, index) in movies" :key="index">  
          <td>{{ movie.id }}</td>  
          <td>{{ movie.name }}</td>  
          <td>{{ movie.director }}</td>  
          <td>{{ movie.writer }}</td>  
          <td>{{ movie.actors }}</td>  
          <td>{{ movie.genre }}</td>  
          <td>{{ movie.region }}</td>  
          <td>{{ movie.language }}</td>  
          <td>{{ movie.releaseDate }}</td>  
          <td>{{ movie.duration }}</td>  
          <td>{{ movie.aliases }}</td>  
        </tr>  
      </tbody>  
    </table>  
  
     "Load More" 按钮 
    <button @click="loadMoreMovies" :disabled="loading || movies.length >= totalMovies">Load More</button>  
  </div>  
</template>   -->
  
<!--
<script>  
export default {  
  data() {  
    return {  
      movies: [], // 存储从后端获取的电影数据  
      loading: false, // 表示是否正在加载数据  
      pn: 0, // 当前页码  
      rn: 10, // 每页项目数  
      totalMovies: 0, // 总电影数  
      backendIpAddress: "60.205.179.8", // 后端 IP 地址  
      backendPort: "6301", // 后端端口号  
    };  
  },  
  methods: {  
    async fetchData() {  
      if (this.loading || this.movies.length >= this.totalMovies) {  
        return;  
      }  
  
      this.loading = true;  
  
      try {  
        const response = await fetch(`http://${this.backendIpAddress}:${this.backendPort}/getMovieData?limit=${this.pn},${this.rn}&orderType=asc`, {  
          method: "GET",  
          headers: {  
            "Content-Type": "application/json",  
          },  
          credentials: "include",  
        });  
  
        const data = await response.json();  
  
        if (data.movies) {  
          this.movies.push(...data.movies);  
          this.pn += this.rn;  
          this.totalMovies = data.total;  
        }  
      } catch (error) {  
        console.error("Error fetching data:", error);  
      } finally {  
        this.loading = false;  
      }  
    },  
    loadMoreMovies() {  
      this.fetchData();  
    },  
  },  
  created() {  
    this.fetchData();  
  },  
};  
</script>   -->
  

  <!-- <template>  
  <main>
  <el-table  
    :data="paginatedMovies"  
    style="display: flex; width:max-content; height:80vh"  
    @scroll="handleScroll"  
  >  
    <el-table-column prop="id" label="ID"></el-table-column>  
    <el-table-column prop="name" label="电影名称"></el-table-column>  
    <el-table-column prop="director" label="导演"></el-table-column>  
    <el-table-column prop="writer" label="编剧"></el-table-column>  
    <el-table-column prop="actors" label="演员"></el-table-column>  
    <el-table-column prop="genre" label="类型"></el-table-column>  
    <el-table-column prop="region" label="地区"></el-table-column>  
    <el-table-column prop="language" label="语言"></el-table-column>  
    <el-table-column prop="releaseDate" label="上映日期"></el-table-column>  
    <el-table-column prop="duration" label="时长"></el-table-column>  
    <el-table-column prop="aliases" label="别名"></el-table-column>  
  </el-table>  
</main>
</template>   -->
  
<!-- <script>  
import request from '@/utils/request.js'
export default {  
  data() {  
    return {  
      movies: [], // 原始数据  
      pn: 1,  
      rn: 10,  
      totalItems: 0,  
      loading: false,  
    };  
  },  
  computed: {  
    paginatedMovies() {  
      const start = (this.pn - 1) * this.rn;  
      const end = this.pn * this.rn;  
      return this.movies.slice(start, end);  
    },  
  },  
  methods: {  
    async fetchData() {  
      if (this.loading) return;  
      this.loading = true;  
      try {  
        const response = await request.get('', {params:{  
        limit: `${this.pn},${this.rn}`,  
        orderType: 'asc'  
      }})
        this.movies = response.data.items;  
        // this.totalItems = response.data.total;  
      } catch (error) {  
        console.error("Error fetching data:", error);  
      } finally {  
        this.loading = false;  
      }  
    },  
    handleScroll({ target }) {  
      const { scrollTop, clientHeight, scrollHeight } = target;  
      if (scrollHeight - scrollTop === clientHeight) {  
        this.pn += 1;  
        this.fetchData();  
      }  
    },  
  },  
  mounted() {  
    this.fetchData();  
  },  
};  
</script>  

<style scoped>  
.item {
    background-color: #fff;
    padding: 10px;
    margin: 5px;
    border: 1px solid #ddd;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 阴影效果 */
    border-radius: 5px; /* 圆角边框 */
    transition: transform 0.3s ease;
    white-space: nowrap; /* 文本不换行 */
    overflow: hidden; /* 隐藏溢出文本 */
    text-overflow: ellipsis; /* 使用省略号表示溢出文本 */
}

.item:hover {
    transform: translateX(10px);
    box-shadow: 2px 4px 6px rgba(0, 0, 0, 0.2); /* 悬停时的阴影效果 */
}

#loadMore {
    margin-top: 10px;
}

table {
    overflow: auto; /* 启用表格滚动 */
}

/* 自定义滚动条样式 */
table::-webkit-scrollbar {
    width: 12px;
}

table::-webkit-scrollbar-track {
    background: #f2f2f2;
}

table::-webkit-scrollbar-thumb {
    background: #007bff;
}

tr:hover {
    background-color: #f0f0f0; /* 悬停时的背景颜色 */
}




body{
  margin: 0;
  padding: 0;
}

</style> -->

<!-- <!DOCTYPE html>
<html> -->
  <template>
<div style="height: 100%; width: 100%;">
      <head>Movie Info</head>
  
      <div>
      <el-table ref="table" :data="movies" max-height="90vh" border stripe style="height: 100%; width: 100%;">
      <el-table-column label="ID" prop="m_id"></el-table-column>
      <el-table-column label="名称" prop="m_title">
        <template #default="scope">
          <a :href="scope.row.m_url" target="_blank">{{ scope.row.m_title }}</a>
        </template>
      </el-table-column>
      <el-table-column label="导演" prop="m_director"></el-table-column>
      <el-table-column label="编剧" prop="m_screenwriter"></el-table-column>
      <el-table-column label="演员" prop="m_starring"></el-table-column>
      <el-table-column label="类型" prop="m_type"></el-table-column>
      <el-table-column label="地区" prop="m_region"></el-table-column>
      <el-table-column label="语言" prop="m_language"></el-table-column>
      <el-table-column label="上映日期" prop="m_releaseDate"></el-table-column>
      <el-table-column label="时长" prop="m_long"></el-table-column>
      <el-table-column label="别名" prop="m_sname"></el-table-column>
    </el-table>

  
      <!-- "Load More" 按钮 -->
      <el-button v-if="showLoadMore" @click="loadMore" style="margin-top: 10px;">Load More</el-button>
    </div>
    </div>
  </template>
  
  <script>
import request from '@/utils/request.js'
  
  export default {
    data() {
      return {
        movies: [],
        pn: 0,
        rn: 10,
        showLoadMore: false, // 控制加载更多按钮显示与隐藏
      };
    },
    methods: {
      fetchData() {
        if (this.pn >= 300) {
          return;
        }
        request.get('', {params:{  
        limit: `${this.pn},${this.rn}`,  
        orderType: 'asc'  
      }}).then(response => {
        if (response.status === 204) {
          this.$message.warning('没有更多数据了');
          return [];
        } else if (response.status === 200) {
          return response.data;
        } else {
          throw new Error(`Request failed with status: ${response.status}`);
        }
      }).then(movieData => {
        if (movieData.length === 0) {
          this.$message.warning('没有更多数据了');
        } else {
          this.movies = [...this.movies, ...movieData];
        }
      }).catch(error => {
        console.error("Error loading data: " + error);
      });
      this.pn += this.rn;
    },
      loadMore() {
        this.fetchData(); // 发起异步数据请求以加载更多内容
        // this.showLoadMore = false;
      }
    },
    mounted() {
      this.fetchData(); // 初始加载数据

      this.showLoadMore = true;
    //   window.addEventListener("scroll", () => {
    //   const scrollY = window.scrollY;
    //   const windowHeight = window.innerHeight;
    //   const contentHeight = document.body.scrollHeight;

    //   if (scrollY + windowHeight >= contentHeight - 50) {
    //     // 当距离底部不足 50px 时触发加载更多
    //     // this.fetchData();
    //     this.showLoadMore = true;
    //   }
    // });
    }
  };
  </script>

<style scoped>
        /* CSS 样式保持不变 */
        /* 增加滑动样式 */
        .el-table__body-wrapper {
  overflow-y: auto; /* 添加垂直滚动条 */
}

        .item {
            background-color: #fff;
            padding: 10px;
            margin: 5px;
            border: 1px solid #ddd;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* 阴影效果 */
            border-radius: 5px; /* 圆角边框 */
            transition: transform 0.3s ease;
            /*white-space: nowrap; !* 文本不换行 *!*/
            overflow: hidden; /* 隐藏溢出文本 */
            text-overflow: ellipsis; /* 使用省略号表示溢出文本 */
        }

        .item:hover {
            transform: translateX(10px);
            box-shadow: 2px 4px 6px rgba(0, 0, 0, 0.2); /* 悬停时的阴影效果 */
        }

        #loadMore {
            display: block;
            margin: 0 auto; /* 居中按钮 */
            padding: 30px 100px;
            background-color: #007bff;
            color: #fff;
            border: none;
            border-radius: 5px; /* 圆角按钮 */
            cursor: pointer;
            font-weight: bold;
            font-size: 16px; /* 设置字体大小为 16px */
            transition: background-color 0.3s ease;
        }

        #loadMore:hover {
            background-color: #0056b3; /* 悬停时的背景颜色变化 */
        }

        th, td {
            border: 1px solid #ddd;
            padding: 8px;
            text-align: left;
        }

        tr:nth-child(even) {
            background-color: #f2f2f2;
        }

        /* 原有的表格样式 */
        table {
            overflow: auto; /* 启用表格滚动 */
            /*border-collapse: collapse;*/
            /*width: 100%;*/
        }

        /* 自定义滚动条样式 */
        table::-webkit-scrollbar {
            width: 12px;
        }

        table::-webkit-scrollbar-track {
            background: #f2f2f2;
        }

        table::-webkit-scrollbar-thumb {
            background: #007bff;
        }

        tr:hover {
            background-color: #a8c598; /* 悬停时的背景颜色 */
        }
    </style>

