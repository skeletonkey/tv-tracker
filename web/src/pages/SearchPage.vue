<template>
  <div>
    <input type="text" v-model="searchString" placeholder="Search...">
    <button @click="search">Search</button>

    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Description</th>
          <th>Image</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in searchResults" :key="item.id">
          <td><img :src="item.thumbnail" alt="Item Image" width="170" height="250"></td>
          <td>{{ item.translations.eng }}</td>
          <td>{{ item.overviews.eng }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchString: '',
      searchResults: []
    };
  },
  methods: {
    search() {
      fetch(`http://localhost:8083/search/${this.searchString}`)
        .then(response => response.json())
        .then(data => {
          this.searchResults = data;
        })
        .catch(error => {
          console.error('Error fetching data:', error);
        });
    }
  }
};
</script>