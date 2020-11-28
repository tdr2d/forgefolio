export default class Image {
  static get toolbox() {
    return {
      icon: '<svg xmlns="http://www.w3.org/2000/svg" width="17" height="15" viewBox="0 0 336 276"><path d="M291 150.242V79c0-18.778-15.222-34-34-34H79c-18.778 0-34 15.222-34 34v42.264l67.179-44.192 80.398 71.614 56.686-29.14L291 150.242zm-.345 51.622l-42.3-30.246-56.3 29.884-80.773-66.925L45 174.187V197c0 18.778 15.222 34 34 34h178c17.126 0 31.295-12.663 33.655-29.136zM79 0h178c43.63 0 79 35.37 79 79v118c0 43.63-35.37 79-79 79H79c-43.63 0-79-35.37-79-79V79C0 35.37 35.37 0 79 0z"/></svg>',
      title: 'Text'
    };
  }

  constructor({ data, config, api, readOnly }) {
    this.api = api;
    this.readOnly = readOnly;
    this.blockIndex = this.api.blocks.getCurrentBlockIndex() + 1; // When block is only constructing, current block points to previous block. So real block index will be +1 after rendering
    console.log(config);
    this.CSS = Object.assign({
      loading: this.api.styles.loader,
      settingsButton: this.api.styles.settingsButton,
      settingsButtonActive: this.api.styles.settingsButtonActive,
    }, config.css);
    console.log(this.CSS);

    this.data = {
      url: data.url || '',
      caption: data.caption || '',
    };

    this.nodes = {
      wrapper: this._make('div', ['f']),
      input: this._make('input', [this.api.styles.input], {"placeholder": "Image url"}),
      submitButton: this._make('button', ['button', 'button--blue'], {}, "Submit"),
      loader: this._make('div', this.CSS.loading),
      imageHolder: this._make('figure', this.CSS.imageHolder),
      image: this._make('img', this.CSS.image),
      caption: this._make('figcaption', this.CSS.imageCaption, {contentEditable: !this.readOnly, innerHTML: this.data.caption || ''}),
    };
    this.nodes.submitButton.onclick = this.onInputSubmit.bind(this);
    this.nodes.caption.dataset.placeholder = 'Enter a caption';
    this._view = this.data.url != '' ? 'visual' : 'input';
    this.settings = [
      {
        name: 'Input View',
        el: this._make('div'),
        view: "input",
        icon: '<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 0 24 24" width="24"><path d="M0 0h24v24H0z" fill="none"/><path d="M21 3.01H3c-1.1 0-2 .9-2 2V9h2V4.99h18v14.03H3V15H1v4.01c0 1.1.9 1.98 2 1.98h18c1.1 0 2-.88 2-1.98v-14c0-1.11-.9-2-2-2zM11 16l4-4-4-4v3H1v2h10v3z"/></svg>',
      },
      {
        name: 'Image View',
        el: this._make('div'),
        view: "visual",
        icon: Image.toolbox.icon
      }
    ];
  }

  render() {
    if (this._view == "input") { this._renderInput(); }
    else if (this._view == "visual") { this._renderImage(); }
    return this.nodes.wrapper;
  }

  _renderInput() {
    this.nodes.wrapper.innerHTML = "";
    this.nodes.wrapper.classList.add('f');
    this.nodes.wrapper.appendChild(this.nodes.input);
    this.nodes.wrapper.appendChild(this.nodes.submitButton);
  }

  _renderImage() {
    this.nodes.wrapper.innerHTML = "";
    this.nodes.wrapper.classList.remove('f');
    this.nodes.wrapper.appendChild(this.nodes.loader);
    if (this.data.url) {
      this.nodes.image.src = this.data.url;
    }

    this.nodes.image.onload = () => {
      this.nodes.wrapper.innerHTML = "";
      this.nodes.imageHolder.appendChild(this.nodes.image);
      this.nodes.imageHolder.appendChild(this.nodes.caption);
      this.nodes.wrapper.appendChild(this.nodes.imageHolder);
    };
    this.nodes.image.onerror = (e) => {
      console.log('Failed to load an image', e);
      this._updateView('input');
    };
  }

  onInputSubmit() {
    this.data.url = this.nodes.input.value;
    this._updateView('visual');
  }

  save() {
    return Object.assign(this.data, {
      caption: this.nodes.caption.innerHTML,
    });
  }

  static get sanitize() {
    return {
      url: {},
      caption: {
        br: true,
      },
    };
  }

  renderSettings() {
    const wrapper = document.createElement('div');
    this.settings.forEach(tune => {
      tune.el.classList.add(this.CSS.settingsButton);
      tune.el.innerHTML = tune.icon;
      tune.el.addEventListener('click', () => {
        this._updateView(tune.view);
      });
      wrapper.appendChild(tune.el);
    });
    return wrapper;
  };

  _make(tagName, classNames = null, attributes = {}, content = null) {
    const el = document.createElement(tagName);
    if (Array.isArray(classNames)) {
      el.classList.add(...classNames);
    } else if (classNames) {
      el.classList.add(classNames);
    }
    for (const attrName in attributes) {
      el[attrName] = attributes[attrName];
    }
    if (content) { el.innerHTML = content; }
    return el;
  }

  _updateView(view) {
    this._view = view;
    this.settings.forEach(t => t.el.classList.toggle(this.CSS.settingsButtonActive, t.view == this._view));
    this.render();
  }
}