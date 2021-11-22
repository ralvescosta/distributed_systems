module.exports = {
  roots: ['<rootDir>'],
  globals: {},
  collectCoverageFrom: [
    '<rootDir>/src/**/*.js'
  ],
  coverageDirectory: 'coverage',
  coverageThreshold: {
    global: {
      branches: 80,
      functions: 80,
      lines: 80,
      statements: -10
    }
  },
  testEnvironment: 'node',
  // transform: {
  //   '.+\\.ts$': 'ts-jest'
  // },
  setupFiles: ['<rootDir>/jest.setup.js']
}