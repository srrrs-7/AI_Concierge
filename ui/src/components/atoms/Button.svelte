<script lang="ts">
  export let name: string = 'Button';
  export let width: string = '100px';
  export let height: string = '30px';
  export let radius: string = '10px';
  export let bgColor: string = 'aqua';
  export let hover: string = 'blue';
  export let fetchData: () => void = function () {
    console.log('log');
  };

  function mouseOver(e: any) {
    e.target.style.backgroundColor = hover;
  }

  function mouseOut(e: any) {
    e.target.style.backgroundColor = bgColor;
  }
</script>

<div>
  <button
    on:click={fetchData}
    on:mouseover={mouseOver}
    on:mouseout={mouseOut}
    on:blur={mouseOut}
    style="width: {width}; height: {height}; border-radius: {radius}; background-color: {bgColor};"
  >
    {name}
  </button>
</div>

<style>
  div {
    margin: 20px;
  }
</style>
