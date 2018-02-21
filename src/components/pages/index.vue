<template>
  <section>
    <button @click="next">もういっちょ！</button>
    <ul>
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
  </section>
</template>

<script>
export default {
  data() {
    return {
      items: [],
      coursor: "-1"
    }
  },
  created() {
    this.getUsers()
  },
  methods: {
    next() {
      this.getUsers()
    },
    getUsers() {
      fetch(`/api/users?coursor=${this.coursor}`)
      .then((res) => { return res.json() })
      .then((res) => {
        this.coursor = res.next_cursor_str
        this.items = [...res.users, ...this.items]
      })
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
</style>
