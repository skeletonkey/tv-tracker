<template>
  <div>
    <input type="text" v-model="searchString" placeholder="Search...">
    <button @click="search">Search</button>

    <div class="card-container" v-if="searchResults.length > 0">
      <div class="card-row" v-for="(row, rowIndex) in groupedSearchResults" :key="rowIndex">
        <div class="card" v-for="item in row" :key="item.id">
          <div class="card-title">{{ item.translations.eng }}</div>
          <img :src="item.thumbnail" alt="Item Image" width="170" height="250"
            :title="item.overviews.eng">
          <div class="card-description">{{ truncatedDescription(item.overviews.eng) }}</div>
        </div>
      </div>
    </div>
    <div v-else>No results found.</div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchString: '',
      searchResults: [],
      cardsPerRow: 3
    };
  },
  props: {
    imageHeight: {
      type: Number,
      default: 250,
    },
    imageWidth: {
      type: Number,
      default: 170,
    }
  },
  mounted() {
    this.calculateCardsPerRow();
    window.addEventListener('resize', this.calculateCardsPerRow);
  },
  beforeUnmount() {
    window.removeEventListener('resize', this.calculateCardsPerRow);
  },
  computed: {
    groupedSearchResults() {
      const result = [];
      for (let i = 0; i < this.searchResults.length; i += this.cardsPerRow) {
        result.push(this.searchResults.slice(i, i + this.cardsPerRow));
      }
      return result;
    },
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
    },
    truncatedDescription(description) {
      return description.length > 50 ? description.substring(0, 50) + '...' : description;
    },
    calculateCardsPerRow() {
      const containerWidth = this.$el.offsetWidth; // Get width of the container element
      const cardWidth = 200; // Max card width with margin
      this.cardsPerRow = Math.max(1, Math.floor(containerWidth / cardWidth));
    },
  }
};
</script>

<style scoped>
.card-container {
  display: flex;
  flex-direction: column;
}

.card-row {
  display: flex;
  flex-wrap: wrap;
  margin-bottom: 20px;
}

.card {
  width: calc(v-bind(imageWidth) + 10px);
  margin: 10px;
  border: 1px solid #ccc;
  padding: 10px;
  border-radius: 5px;
  box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.1);
}

.card-title {
  text-align: center;
  font-weight: bold;
  margin-bottom: 5px;
}

.card-description {
  font-size: 0.9em;
  color: #555;
}

.card img {
  display: block;
  margin: 0 auto;
  max-width: 100%;
  height: auto;
}
</style>