<script lang="ts">
  export let width: string = '30px';
  export let height: string = '20px';
  export let placeholder: string = 'input';

  $: value = '';
</script>

<div>
  <input
    type="text"
    {placeholder}
    bind:value
    style="width:{width}; height:{height};"
  />
</div>

<style>
  div {
    margin: 20px;
  }
</style>
