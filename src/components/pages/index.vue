<template>
  <section>
    <div class="spinner" v-show="isLoading">
      <pacman-loader></pacman-loader>
    </div>
    <ul v-if="items.length > 0">
      <li v-for="(item, index) in items" :key="index" class="card-item">
        <div class="container">
          <div class="profile">
            <div class="img">
              <img :src="item.profile_image_url_https" alt="icon">
            </div>
            <p class="name">{{item.name}}@{{item.screen_name}}</p>
          </div>
          <p class="description">{{item.description}}</p>
        </div>
        <button class="button" @click="block" :data-id="item.screen_name">ブロックする</button>
      </li>
    </ul>
    <div class="no-result" v-else>
      <h1>{{errorMessage}}</h1>
    </div>
  </section>
</template>

<script>
import axios from 'axios'
import PacmanLoader from 'vue-spinner/src/PacmanLoader.vue'

export default {
  components: {
    PacmanLoader
  },
  data() {
    return {
      isLoading: true,
      items: [],
      errorMessage: ""
    }
  },
  created() {
    this.getUsers()
  },
  methods: {
    next() {
      this.getUsers()
    },
    async healthcheck(type) {
      let message = ""

      const res = await axios.get("/api/healthcheck")

      const followersRemaining = res.data.resources.followers["/followers/ids"].remaining
      const blockRemaining = res.data.resources.blocks["/blocks/list"].remaining

      switch(type) {
        case "blocks":
          message = blockRemaining == 0 ? "API LIMITです" : "ブロックに失敗しました。"
          break;
        case "followers":
          message = followersRemaining == 0 ? "API LIMITです" : "怪しいアカウントが見つかりませんでした。"
          break;
        default:
          message = followersRemaining == 0 ? "API LIMITです" : "怪しいアカウントが見つかりませんでした。"
      }

      this.errorMessage = message
    },
    async getUsers() {
      const res = await axios.get("/api/users")
      if(res.status == 200) {
        this.isLoading = false
        if(res.data.users.length == 0) this.healthcheck()
        this.items = res.data.users
      }
    },
    block(e) {
      const id = e.target.dataset.id
      fetch(`/api/users/${id}/block`, {
        method: "POST"
      })
      .then((res) => {
        if(res.status == 200) {
          alert("ブロックしました！")
          this.items.some((v, i) => {
            if(v.screen_name == id) this.items.splice(i, 1)
          })
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
body, html {
  width: 100%;
  height: 100%;
  background-color: #fafafa;
}
* {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}

ul {
  width: 100%;
  list-style: none;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  .card-item {
    width: 32%;
    box-shadow: rgba(0, 0, 0, 0.12) 0px 1px 6px, rgba(0, 0, 0, 0.12) 0px 1px 4px;
    background-color: white;
    margin: 10px 0;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    display: flex;
    .container {
      width: 100%;
      .profile {
        display: flex;
        align-items: center;
        font-size: 0.7rem;
        font-weight: 600;
        .img {
          width: 50px;
          margin-right: 5px;
          img {
            display: block;
            width: 100%;
          }
        }
      }
      .description {
        font-size: 0.7rem;
        padding: 10px;
      }
    }
    .button {
      display: block;
      width: 100%;
      border: 0;
      outline: none;
      cursor: pointer;
      background-color: rgb(237, 60, 60);
      color: white;
      font-weight: 600;
      padding: 10px 0;
    }
  }
}
.spinner {
  background: white;
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
