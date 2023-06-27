<script lang="ts">
  export let tagName: string = 'tag';
  export let width: string = '20px';
  export let height: string = '20px';
  export let color: string = 'black';
  export let radius: string = '10px';
  export let clickFunction: () => void = () => {};
</script>

<button
  style="width: {width}; height: {height}; color: {color}; border-radius: {radius}"
  on:click={clickFunction}
>
  {tagName}
</button>
