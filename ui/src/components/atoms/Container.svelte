<script lang="ts">
  export let width = "20px";
  export let height = "20px";
  export let bgColor = "currentColor";
</script>

<div class="container" style="width: {width}; height: {height}; background-color: {bgColor}">
  <slot />
</div>
