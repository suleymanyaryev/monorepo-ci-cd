import { createApp } from "vue";
import App from "./App.vue";
import "@/assets/css/tailwind.css";

const app = createApp(App);
app.directive("cy", (el, binding) => {
    if (binding.value) {
        el.setAttribute("data-cy", binding.value);
    }
});
app.mount("#app");
