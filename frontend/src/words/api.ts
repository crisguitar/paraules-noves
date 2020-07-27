import axios from 'axios'
import config from '../config'
import Word from './word'

export default {
  createWord: (word: Word) =>
    axios.post(`${config.host}/words`, word),
  allWords: async () => {
    const response = await axios.get(`${config.host}/words`)
    return response.data
  }
}
