module.exports = {
  // content: [],
  mode: 'jit',
  theme: {
    extend: {},
  },
  plugins: [],
  purge: {
    enabled: true,
    content: ['./apps/cms/**/*.{html,ts}'],
  },
};
