import api from './api'
import axios from 'axios'
import config from '../config'
import Word from './word'

jest.mock('axios', () => ({
  post: jest.fn(),
  get: jest.fn()
}))

jest.mock('../config', () => ({
  host: 'http://1.com'
}))

describe('api', () => {
  describe('create word', () => {
    let mockPost: jest.Mock
    let word: Word
    beforeEach(() => {
      mockPost = axios.post as jest.Mock
      word = {
        word: 'some word',
        meaning: 'some meaning'
      }
    })

    it('should resolve when call is successful', async () => {
      mockPost.mockResolvedValue('success')

      const result = api.createWord(word)

      await expect(result).resolves.toEqual('success')
      expect(mockPost).toHaveBeenCalledWith('http://1.com/words', word)
    })

    it('should reject when call is not successful', async () => {
      mockPost.mockRejectedValue('failure')

      const result = api.createWord(word)

      await expect(result).rejects.toEqual('failure')
    })
  })

  describe('get all words', () => {
    let mockGet: jest.Mock
    beforeEach(() => {
      mockGet = axios.get as jest.Mock
    })

    it('should resolve words from server', async () => {
      const expectedResult = [{}, {}]
      mockGet.mockResolvedValue({data: expectedResult})

      const result = api.allWords()

      await expect(result).resolves.toEqual(expectedResult)
      expect(mockGet).toHaveBeenCalledWith('http://1.com/words')
    });

    it('should reject when call is not successful', async () => {
      mockGet.mockRejectedValue('failure')

      const result = api.allWords()

      await expect(result).rejects.toEqual('failure')
    });
  });
})
