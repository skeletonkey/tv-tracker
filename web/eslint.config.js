import pluginVue from 'eslint-plugin-vue'

export default [
  {
    ignores: ['dist/**']
  },
  ...pluginVue.configs['flat/essential'],
  {
    rules: {
      'no-unused-vars': 'warn',
      'no-console': 'warn'
    }
  }
]
