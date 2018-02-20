<template>
  <section>
    <ul>
      <li v-for="item in items">
        <p>{{item.description}}</p>
        <!-- <div class="img">
          <img :src="item.profile_image_url_https" alt="icon">
        </div>
        <p>{{item.name}}@{{item.screen_name}}</p>
        <p>{{itme.description}}</p> -->
      </li>
    </ul>
  </section>
</template>

<script>
export default {
  data() {
    return {
      items: []
    }
  },
  created() {
    fetch("/api/users")
    .then((res) => { return res.json() })
    .then((res) => {
      console.log(res)
      this.items = res
    })
  }
}
</script>
