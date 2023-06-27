<script>
  let items = [
      { name: '商品1', price: 100 },
      { name: '商品2', price: 200 },
      { name: '商品3', price: 300 }
  ];
</script>

<table>
  <thead>
      <tr>
          <th>商品名</th>
          <th>価格</th>
      </tr>
  </thead>
  <tbody>
      {#each items as item}
          <tr>
              <td>{item.name}</td>
              <td>{item.price}</td>
          </tr>
      {/each}
  </tbody>
</table>
