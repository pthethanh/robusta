/**
 * Build styles
 */
require('./index.css').toString();

// import ToolboxIcon from './toolbox.svg';

/**
 * CodeEx for Editor.js
 *
 * @author CodeX (team@ifmo.su)
 * @copyright CodeX 2018
 * @license The MIT License (MIT)
 * @version 2.0.0
 */

class CodeEx {
  /**
   * Allow to press Enter inside the CodeEx textarea
   * @returns {boolean}
   * @public
   */
  static get enableLineBreaks() {
    return true;
  }

  /**
   * @typedef {Object} CodeData — plugin saved data
   * @param {String} sample - previously saved plugin code
   * @param {String} code - previously saved plugin code
   */

  /**
   * Render plugin`s main Element and fill it with saved data
   *
   * @param {CodeData} data — previously saved plugin code
   * @param {Object} config - user config for Tool
   * @param {Object} api - Editor.js API
   */
  constructor({ data, config, api }) {
    this.api = api;

    this.samplePlaceholder = config.placeholder || CodeEx.DEFAULT_SAMPLE_PLACEHOLDER;
    this.codePlaceholder = config.placeholder || CodeEx.DEFAULT_CODE_PLACEHOLDER;

    this.CSS = {
      baseClass: this.api.styles.block,
      input: this.api.styles.input,
      wrapper: 'ce-code',
      textarea: 'ce-code__textarea'
    };

    this.nodes = {
      holder: null,
      sampleTextarea: null,
      codeTextarea: null
    };

    this.data = {
      sample: data.sample || '',
      code: data.code || ''
    };

    this.nodes.holder = this.drawView();
  }

  /**
   * Create Tool's view
   * @return {HTMLElement}
   * @private
   */
  drawView() {
    let wrapper = document.createElement('div');

    // sample box
    let sampleWrapper = document.createElement('div');

    let sampleTextarea = document.createElement('textarea');

    sampleTextarea.id = 'sample-textarea';
    sampleWrapper.classList.add(this.CSS.baseClass, this.CSS.wrapper);
    sampleTextarea.classList.add(this.CSS.textarea, this.CSS.input);
    sampleTextarea.textContent = this.data.sample;
    sampleTextarea.placeholder = this.samplePlaceholder;
    sampleWrapper.appendChild(sampleTextarea);
    this.nodes.sampleTextarea = sampleTextarea;

    // code box
    let codeWrapper = document.createElement('div');

    let codeTextarea = document.createElement('textarea');

    codeTextarea.id = 'code-textarea';
    codeWrapper.classList.add(this.CSS.baseClass, this.CSS.wrapper);
    codeTextarea.classList.add(this.CSS.textarea, this.CSS.input);
    codeTextarea.textContent = this.data.code;
    codeTextarea.placeholder = this.codePlaceholder;
    codeWrapper.appendChild(codeTextarea);
    this.nodes.codeTextarea = codeTextarea;

    wrapper.appendChild(sampleWrapper);
    wrapper.appendChild(codeWrapper);

    return wrapper;
  }

  /**
   * Return Tool's view
   * @returns {HTMLDivElement} this.nodes.holder - Code's wrapper
   * @public
   */
  render() {
    return this.nodes.holder;
  }

  /**
   * Extract Tool's data from the view
   * @param {HTMLDivElement} codeWrapper - CodeEx's wrapper, containing textarea with code
   * @returns {CodeData} - saved plugin code
   * @public
   */
  save(wrapper) {
    return {
      code: wrapper.querySelector('#code-textarea').value,
      sample: wrapper.querySelector('#sample-textarea').value
    };
  }

  /**
   * onPaste callback fired from Editor`s core
   * @param {PasteEvent} event - event with pasted content
   */
  onPaste(event) {
    const content = event.detail.data;

    this.data = {
      code: content.textContent
    };
  }

  /**
   * Returns Tool`s data from private property
   * @return {*}
   */
  get data() {
    return this._data;
  }

  /**
   * Set Tool`s data to private property and update view
   * @param {CodeData} data
   */
  set data(data) {
    this._data = data;

    if (this.nodes.textarea) {
      this.nodes.codeTextarea.textContent = data.code;
      this.nodes.sampleTextarea.textContent = data.sample;
    }
  }

  /**
   * Get Tool toolbox settings
   * icon - Tool icon's SVG
   * title - title to show in toolbox
   *
   * @return {{icon: string, title: string}}
   */
  static get toolbox() {
    return {
      // icon: ToolboxIcon,
      icon: '<svg height="16" viewBox="0 -61 512 512" width="16" xmlns="http://www.w3.org/2000/svg"><path d="m497 0h-482c-8.285156 0-15 6.714844-15 15v360c0 8.285156 6.714844 15 15 15h482c8.285156 0 15-6.714844 15-15v-360c0-8.285156-6.714844-15-15-15zm-15 60h-361v-30h361zm-391-30v30h-61v-30zm-61 330v-270h452v270zm0 0"/><path d="m356.605469154.394531c-5.855469-5.859375-15.355469-5.859375-21.210938 0-5.859375 5.859375-5.859375 15.355469 0 21.210938l49.390625 49.394531-49.390625 49.394531c-5.859375 5.859375-5.859375 15.355469 0 21.210938 5.855469 5.859375 15.355469 5.859375 21.210938 0l60-60c5.859375-5.855469 5.859375-15.351563 0-21.210938zm0 0"/><path d="m176.605469 154.394531c-5.855469-5.859375-15.355469-5.859375-21.210938 0l-60 60c-5.859375 5.859375-5.859375 15.355469 0 21.210938l60 60c5.855469 5.859375 15.351563 5.859375 21.210938 0 5.859375-5.855469 5.859375-15.351563 0-21.210938l-49.390625-49.394531 49.390625-49.394531c5.859375-5.855469 5.859375-15.355469 0-21.210938zm0 0"/><path d="m290.742188 120.769531c-7.859376-2.617187-16.351563 1.628907-18.972657 9.488281l-60 180c-2.621093 7.859376 1.628907 16.351563 9.488281 18.972657 7.859376 2.621093 16.351563-1.628907 18.972657-9.488281l60-180c2.621093-7.859376-1.628907-16.351563-9.488281-18.972657zm0 0"/></svg>',
      title: 'Code With Example'
    };
  }

  /**
   * Default placeholder for CodeEx's textarea
   *
   * @public
   * @returns {string}
   */
  static get DEFAULT_SAMPLE_PLACEHOLDER() {
    return 'Enter sample code';
  }

  /**
   * Default placeholder for warning message
   *
   * @public
   * @returns {string}
   */
  static get DEFAULT_CODE_PLACEHOLDER() {
    return 'Enter runnable code';
  }

  /**
   *  Used by Editor.js paste handling API.
   *  Provides configuration to handle CODE tag.
   *
   * @static
   * @return {{tags: string[]}}
   */
  static get pasteConfig() {
    return {
      tags: [ 'pre' ]
    };
  }
}

module.exports = CodeEx;
