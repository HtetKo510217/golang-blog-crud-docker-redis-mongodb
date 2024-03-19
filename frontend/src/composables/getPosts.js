import { ref } from 'vue'
import axios from 'axios'


let getPosts = ()=> {
    let posts = ref([]);
    let load = async ()=> {
      const response = await axios.get('http://localhost:3000/api/v1/posts')
      posts.value = response.data
    }

    return {posts,load}
}

export default getPosts;