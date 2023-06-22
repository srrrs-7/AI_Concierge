<script lang="ts">
  export let label: string = 'label';
</script>

<div>
  <label for="radio"> {label} </label>
  <input type="radio" id="radio" checked />
</div>

<style>
  div {
    margin: 20px;
  }
</style>
