import service from './service';

class NewWordForm extends HTMLElement {
  constructor() {
    super();

    this.innerHTML = this.template

    this.addEventListener('submit', this.handleSubmit);

    this.handleSubmit = this.handleSubmit.bind(this)
  }

  async handleSubmit(e) {
    e.preventDefault()
    const word = document.getElementById('form-new-word');
    const meaning = document.getElementById('form-new-meaning');

    if (this.isValid(word.value, meaning.value)) {
      await service.addWord(word.value, meaning.value)
      service.loadWords()
      word.value = ''
      meaning.value = ''
    }
  }

  isValid(word, meaning) {
    return word !== '' && meaning !== ''
  }

  get template() {
    return `
      <form>
        <div class="field">
          <div class="control">
            <input id="form-new-word" name="word" class="input" type="text" placeholder="paraula">
          </div>
        </div>
        <div class="field">
          <div class="control">
            <textarea id="form-new-meaning" name="meaning" class="textarea" placeholder="significat"></textarea>
          </div>
        </div>
        <div class="field is-grouped">
          <div class="control">
            <button type="submit" class="button is-link has-background-primary">Desa</button>
          </div>
        </div>
      </form>
    `;
  }
}

customElements.define('word-form', NewWordForm);
