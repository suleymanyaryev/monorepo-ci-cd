/* eslint-env node */
require("@rushstack/eslint-patch/modern-module-resolution");

module.exports = {
    "root": true,
    "globals": {
        "process": true,
    },
    "extends": [
        "plugin:vue/vue3-essential",
        "eslint:recommended",
        "@vue/eslint-config-prettier",
    ],
    "env": {
        "vue/setup-compiler-macros": true,
        "browser": true,
        "amd": true,
        "node": true,
    },
    "overrides": [
        {
            "files": ["cypress/integration/**.spec.{js,ts,jsx,tsx}"],
            "extends": ["plugin:cypress/recommended"],
        },
    ],
};
