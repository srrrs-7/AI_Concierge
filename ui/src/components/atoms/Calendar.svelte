<script>
  import { onMount } from 'svelte';

  let currentDate = new Date();
  let currentMonth = currentDate.getMonth();
  let currentYear = currentDate.getFullYear();
  let daysInMonth = new Date(currentYear, currentMonth + 1, 0).getDate();
  let firstDayOfWeek = new Date(currentYear, currentMonth, 1).getDay();
  let calendarDays = [];

  onMount(() => {
    generateCalendarDays();
  });

  function generateCalendarDays() {
    calendarDays = [];

    // Add empty cells for the days before the first day of the month
    for (let i = 0; i < firstDayOfWeek; i++) {
      calendarDays.push('');
    }

    // Add cells for the days in the month
    for (let i = 1; i <= daysInMonth; i++) {
      calendarDays.push(i);
    }
  }

  function previousMonth() {
    currentMonth--;
    if (currentMonth < 0) {
      currentYear--;
      currentMonth = 11;
    }
    updateCalendar();
  }

  function nextMonth() {
    currentMonth++;
    if (currentMonth > 11) {
      currentYear++;
      currentMonth = 0;
    }
    updateCalendar();
  }

  function updateCalendar() {
    daysInMonth = new Date(currentYear, currentMonth + 1, 0).getDate();
    firstDayOfWeek = new Date(currentYear, currentMonth, 1).getDay();
    generateCalendarDays();
  }
</script>

<div>
  <h2>{currentYear}年 {currentMonth + 1}月</h2>
  <button on:click={previousMonth}>前の月</button>
  <button on:click={nextMonth}>次の月</button>
  <table>
    <thead>
      <tr>
        <th>日</th>
        <th>月</th>
        <th>火</th>
        <th>水</th>
        <th>木</th>
        <th>金</th>
        <th>土</th>
      </tr>
    </thead>
    <tbody>
      {#each calendarDays as day, i}
        {#if i % 7 === 0}
          <tr />
        {/if}
        <td>{day}</td>
        {#if i % 7 === 6}
          <tr />
        {/if}
      {/each}
    </tbody>
  </table>
</div>
