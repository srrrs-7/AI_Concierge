<script lang="ts">
  export let width: string = '100px';
  export let height: string = '100px';
  export let radius: string = '100px';
  export let bg_color: string = 'aqua';
</script>

<div
  style="width: {width}; height: {height}; border-radius: {radius}; background-color: {bg_color};"
>
  <slot />>
</div>

<style>
  div {
    margin: 20px;
  }
</style>
