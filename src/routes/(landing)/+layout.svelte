<script>
  import { useUser } from "$lib";
  import Logo from "$lib/components/Logo.svelte";

  const user = useUser();
</script>

<div
  id="header"
  class="w-full flex justify-between p4 items-center border-b dark:border-dark-50 text-sm"
>
  <Logo />
  <div class="flex items-center space-x-12">
    <a href="/" class="hover:underline">Inicio</a>
    <a href="/" class="hover:underline">Nuestros archivos</a>
    <a href="/" class="hover:underline">Acerca de nosotros</a>
    <a
      href="/app"
      class="bg-blue-500 p4 rounded shadow-lg transition-200 hover:bg-blue-400 text-white anim"
      style:--anim-d="1000ms">{$user ? "Ir a la app" : "Iniciar sesión"}</a
    >
  </div>
</div>

<slot />
