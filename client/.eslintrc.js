module.exports = {
    root: true, // Ensures this config is used as the root configuration
    parser: 'vue-eslint-parser', // Required to parse .vue files
    parserOptions: {
        // Specify a JS parser for <script> blocks
        parser: '@babel/eslint-parser', // Babel parser for ES6+ syntax
        ecmaVersion: 2020, // Allows modern ECMAScript features
        sourceType: 'module', // Use ES Modules
        ecmaFeatures: {
            jsx: true, // Optional only if JSX is used
        },
    },
    extends: [
        'plugin:vue/vue3-essential', // Essential rules for Vue 3
        'eslint:recommended', // ESLint's recommended rules
    ],
    plugins: [
        'vue', // Loads eslint-plugin-vue
    ],
    env: {
        browser: true, // Define browser global variables
        node: true, // Define Node.js global variables
        es2021: true, // Enable modern ES environment
    },
    rules: {
        // Define custom rules or override defaults
        'vue/multi-word-component-names': 'off', // Disable multi-word component name requirement
    },
};