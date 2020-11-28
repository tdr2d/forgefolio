import {CreateElement} from '../utils';

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
    this.CSS = Object.assign({
      loading: this.api.styles.loader,
      settingsButton: this.api.styles.settingsButton,
      settingsButtonActive: this.api.styles.settingsButtonActive,
    }, config.css);
    
    this.data = {
      url: data.url || '',
      caption: data.caption || '',
    };

    this.nodes = {
      wrapper: CreateElement('div', ['f']),
      input: CreateElement('input', [this.api.styles.input], {"placeholder": "Image url"}),
      submitButton: CreateElement('button', ['button', 'button--blue'], {}, "Submit"),
      loader: CreateElement('div', this.CSS.loading),
      imageHolder: CreateElement('figure', this.CSS.imageHolder),
      image: CreateElement('img', this.CSS.image, {crossOrigin: "anonymous"}),
      caption: CreateElement('figcaption', this.CSS.imageCaption, {contentEditable: true, innerHTML: this.data.caption || ''}),
    };
    this.nodes.submitButton.onclick = this.onInputSubmit.bind(this);
    this.nodes.caption.dataset.placeholder = 'Enter a caption';
    this._view = this.data.url != '' ? 'visual' : 'input';
    this.settings = [
      {
        name: 'Input View',
        el: CreateElement('div'),
        view: "input",
        icon: '<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 0 24 24" width="24"><path d="M0 0h24v24H0z" fill="none"/><path d="M21 3.01H3c-1.1 0-2 .9-2 2V9h2V4.99h18v14.03H3V15H1v4.01c0 1.1.9 1.98 2 1.98h18c1.1 0 2-.88 2-1.98v-14c0-1.11-.9-2-2-2zM11 16l4-4-4-4v3H1v2h10v3z"/></svg>',
      },
      {
        name: 'Image View',
        el: CreateElement('div'),
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
      this.nodes.imageHolder.innerHTML = "";
      this.nodes.imageHolder.appendChild(this.nodes.image);
      this.nodes.imageHolder.appendChild(this.nodes.caption);
      this.nodes.wrapper.appendChild(this.nodes.imageHolder);
      this._computeImageData();
    };
    this.nodes.image.onerror = (e) => {
      console.log('Failed to load an image', e);
      this._updateView('input');
    };
  }

  _computeImageData() {
    var canvas = document.createElement('CANVAS');
    var ctx = canvas.getContext('2d');
    canvas.height = 2;
    canvas.width = 3;
    ctx.drawImage(this.nodes.image, 0, 0, canvas.width, canvas.height);
    this.data.pixels = canvas.toDataURL();
    this.data.hwratio = this.nodes.image.height / this.nodes.image.width;
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

  _updateView(view) {
    this._view = view;
    this.settings.forEach(t => t.el.classList.toggle(this.CSS.settingsButtonActive, t.view == this._view));
    this.render();
  }
}