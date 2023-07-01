<script>
  let isOpen = false;

  function toggleSidebar() {
    isOpen = !isOpen;
  }
</script>

<div class="sidebar {isOpen ? 'open' : ''}">
  <ul>
    <li>メニュー1</li>
    <li>メニュー2</li>
    <li>メニュー3</li>
  </ul>
</div>

<div class="content">
  <button on:click={toggleSidebar}>サイドバーを開く/閉じる</button>
  <p>コンテンツがここに表示されます。</p>
</div>

<style>
  .sidebar {
    width: 200px;
    background-color: lightgray;
    height: 100vh;
    position: fixed;
    top: 0;
    left: 0;
    transform: translateX(-100%);
    transition: transform 0.3s ease-in-out;
    z-index: 999;
  }

  .sidebar.open {
    transform: translateX(0);
  }

  .content {
    margin-left: 200px;
    padding: 20px;
  }
</style>
