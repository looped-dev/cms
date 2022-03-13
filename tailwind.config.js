module.exports = {
  content: ['./apps/cms/**/*.{html,ts}', './libs/**/*.{html,ts}'],
  mode: 'jit',
  theme: {
    extend: {
      colors: {
        primary: '#032D5F',
        secondary: 'rgb(249 115 22 / var(--tw-text-opacity))',
      },
    },
  },
  plugins: [],
};
