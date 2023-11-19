<script>
  import { page } from "$app/stores";
  import Logo from "$lib/components/Logo.svelte";
  import { expoOut } from "svelte/easing";
  import { fade } from "svelte/transition";

  export let data;
</script>

<div class="body" in:fade={{ duration: 400, easing: expoOut }}>
  <div class="w-full flex justify-between p4 items-center absolute z-200">
    <Logo theme="dark" />
    <div class="flex items-center space-x-12">
      {#if !data.admin}
        <a
          href={$page.route.id == "/(auth)/register" ? "/login" : "/register"}
          class="bg-blue-500 p4 rounded shadow-lg transition-200 hover:bg-blue-400 text-white"
        >
          {$page.route.id == "/(auth)/register"
            ? "Iniciar sesión"
            : "Registrarse"}</a
        >
      {/if}
    </div>
  </div>
  <section class="form-main w-full">
    <div class="form-content">
      <div class="circle1" />
      <div class="circle2" />
      <div class="circle3" />
      <div class="box flex w-full">
        <slot />
      </div>
    </div>
  </section>
</div>

<style>
  :root {
    --blue: hsl(220, 85%, 57%);
    --violet: hsl(272, 92%, 32%);
    --text-light-gray: hsl(40, 0%, 90%);
    --tex-white: hsl(0, 0%, 90%);
    --body-bg-color: hsl(200, 54%, 12%);
    --glass-bg-color: hsla(0, 0%, 100%, 0.05);
    --border-color: hsla(0, 0%, 100%, 0.25%);
    --blur: blur(10px);
    --button-hover-color: hsla(0, 0% 0%, 0.3);
  }

  .body {
    background-color: var(--body-bg-color);
    color: var(--text-light-gray);
    overflow: hidden;
    height: 100%;
    min-height: 100vh;
    overflow-y: auto;
  }

  .body::after,
  .body::before {
    content: "";
    position: fixed;
    height: 400px;
    width: 400px;
    border-radius: 50%;
    z-index: -1;
    filter: blur(150px);
    opacity: 0.5;
    z-index: 100;
  }

  .body::before {
    background-color: var(--blue);
    left: 0;
    bottom: 0;
    transform: translate(-50%);
    z-index: 100;
  }

  .body::after {
    background-color: var(--violet);
    right: 0;
    top: 0;
    transform: translate(50%);
  }

  .form-main {
    min-height: 100vh;
    padding: 60px 15px;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
  }

  .form-content {
    width: 100%;
    max-width: 450px;
    position: relative;
    z-index: 1;
  }

  .form-content .circle1,
  .form-content .circle2,
  .form-content .circle3 {
    position: absolute;
    background: linear-gradient(to right, var(--blue), var(--violet));
    border-radius: 50%;
    z-index: -1;
  }
  .form-content .circle1 {
    height: 120px;
    width: 120px;
    left: 0;
    top: 10%;
    transform: translate(-50%);
  }

  .form-content .circle2 {
    height: 80px;
    width: 90px;
    right: 0;
    bottom: 0;
    transform: translate(20%, 20%);
  }

  .form-content .circle3 {
    height: 50px;
    width: 50px;
    right: 10%;
    top: 10%;
  }
  .form-content .box {
    --uno: border border-gray-100;
    border-color: white;
    padding: 40px 50px; /*de ambos lados*/
    border-radius: 20px;
    backdrop-filter: var(--blur);
  }
</style>
