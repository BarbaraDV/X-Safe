<script lang="ts">
  import { base } from "$app/paths";
  import { page } from "$app/stores";
  import type { RecordModel } from "pocketbase";
  let categories: RecordModel[] = [];
  $: categories = $page.data.categories || [];
</script>

<div class="flex flex-col w-full">
  <div class="flex w-full p4 justify-between">
    <div class="flex items-center space-x-2 text-xl">
      <un-i-octicon-issue-opened />
      <h3 class="font-bold">Incidencias</h3>
    </div>
  </div>
  <div class="pl-12 pr-4 flex-col space-y-2 text-sm mb4 items-start">
    <a
      class="bg-blue-500 p2 rounded transition-200 hover:bg-blue-400 text-white flex-inline items-center space-x-2"
      href="{base}/incidents/new"
    >
      <un-i-carbon-add />
      <span class="text-xs">Crear incidencia</span>
    </a>
  </div>
  <div class="pl-12 pr-4 flex-col space-y-2 text-sm">
    <div class="flex w-full justify-between pb-2">
      <div class="flex items-center space-x-2">
        <un-i-octicon-issue-draft-16 />
        <a href="{base}/incidents">Todas</a>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <div class="flex items-center space-x-2">
        <un-i-ph-fire-simple-fill class="text-red" />
        <a href="{base}/incidents?type=critic">Riesgo crítico</a>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <div class="flex items-center space-x-2">
        <un-i-material-symbols-priority-high class="text-orange" />
        <a href="{base}/incidents?type=high">Riesgo alto</a>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <div class="flex items-center space-x-2">
        <un-i-mdi-speedometer-medium class="text-purple" />
        <a href="{base}/incidents?type=medium">Riesgo medio</a>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <div class="flex items-center space-x-2">
        <un-i-material-symbols-humidity-high class="text-blue" />
        <a href="{base}/incidents?type=low">Riesgo bajo</a>
      </div>
    </div>
  </div>
  <div class="flex w-full p4 justify-between">
    <div class="flex items-center space-x-2 text-xl">
      <un-i-lucide-bug />
      <h3 class="font-bold">Vulnerabilidades</h3>
    </div>
  </div>
  <div class="pl-12 pr-4 flex-col space-y-2 text-sm mb4 items-start">
    <a
      class="bg-blue-500 p2 rounded transition-200 hover:bg-blue-400 text-white flex-inline items-center space-x-2"
      href="{base}/vulnerabilities/new"
    >
      <un-i-carbon-add />
      <span class="text-xs">Crear vulnerabilidad</span>
    </a>
  </div>
  <div class="pl-12 pr-4 flex-col space-y-2 text-sm">
    <div class="flex w-full justify-between pb-2">
      <div class="flex items-center space-x-2">
        <un-i-octicon-issue-draft-16 />
        <a href="{base}/vulnerabilities">Todas</a>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <div class="flex items-center space-x-2">
        <un-i-ph-fire-simple-fill class="text-red" />
        <a href="{base}/vulnerabilities?type=critic">Riesgo crítico</a>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <div class="flex items-center space-x-2">
        <un-i-material-symbols-priority-high class="text-orange" />
        <a href="{base}/vulnerabilities?type=high">Riesgo alto</a>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <div class="flex items-center space-x-2">
        <un-i-mdi-speedometer-medium class="text-purple" />
        <a href="{base}/vulnerabilities?type=medium">Riesgo medio</a>
      </div>
    </div>
    <div class="flex w-full justify-between pb-2">
      <div class="flex items-center space-x-2">
        <un-i-material-symbols-humidity-high class="text-blue" />
        <a href="{base}/vulnerabilities?type=low">Riesgo bajo</a>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <div class="flex items-center space-x-2">
        <un-i-octicon-issue-opened-16 class="text-red" />
        <a href="{base}/vulnerabilities?status=pending">Pendientes</a>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <div class="flex items-center space-x-2">
        <un-i-icon-park-outline-preview-open class="text-yellow" />
        <a href="{base}/vulnerabilities?status=reviewing">En revisión</a>
      </div>
    </div>
    <div class="flex w-full justify-between">
      <div class="flex items-center space-x-2">
        <un-i-octicon-issue-closed-16 class="text-green" />
        <a href="{base}/vulnerabilities?status=solved">Resueltas</a>
      </div>
    </div>
  </div>
  {#if categories.length}
    <div class="flex w-full px4 pl6 pt4 justify-between">
      <div class="flex items-center space-x-2">
        <un-i-carbon-category />
        <h3 class="font-bold">Categorías</h3>
      </div>
    </div>
    <div class="pl-12 pr-4 flex-col space-y-2 text-sm pt-4">
      <!-- <div class="flex w-full justify-between pb-2"> -->
      <!--   <div class="flex items-center space-x-2"> -->
      <!--     <un-i-octicon-issue-draft-16 /> -->
      <!--     <a href="{base}/vulnerabilities?category=none">Sin categoría</a> -->
      <!--   </div> -->
      <!-- </div> -->
      {#each categories as c}
        <div class="flex w-full justify-between pb-2">
          <div class="flex items-center space-x-2">
            <un-i-material-symbols-category-outline />
            <a href="{base}/vulnerabilities?category={c.id}">{c.name}</a>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<style>
  a:hover {
    text-decoration: underline;
  }
</style>
