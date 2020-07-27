import api from './api'
import Word from './word'

export default {
  loadWords: async () => {
    const words = await api.allWords()
    const allWords = words.map((w: Word) => `
      <word-card class="box">
        <span slot="word">${w.word}</span>
        <span slot="meaning">${w.meaning}</span>
      </word-card>
    `).reduce((prev: string, curr: string) => prev + curr, "")
    document.getElementById("practice").innerHTML = allWords
  },
  addWord: (word: string, meaning: string) => api.createWord({word, meaning})
}
