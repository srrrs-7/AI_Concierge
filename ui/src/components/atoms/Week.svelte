<script lang="ts">
  export let week: number = 1;
  export let width: string = '20px';
  export let height: string = '20px';
  export let color: string = 'black';
</script>

<div style="width: {width}; height: {height}; color: {color};">
  {week}
</div>
