class WordCard extends HTMLElement {
  constructor() {
    super();

    const template = document.getElementById('word-card');
    const templateContent = template.content;

    const shadowRoot = this.attachShadow({mode: 'open'});

    shadowRoot.appendChild(templateContent.cloneNode(true));
  }
}

customElements.define('word-card', WordCard)
