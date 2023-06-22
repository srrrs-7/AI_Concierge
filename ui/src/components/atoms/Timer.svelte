<script lang="ts">
  import moment from 'moment';

  export let hoursSize: string = '20px';
  export let minutesSize: string = '20px';
  export let secondsSize: string = '20px';
  export let hoursColor: string = 'black';
  export let minutesColor: string = 'black';
  export let secondsColor: string = 'black';

  let hours = moment().format('HH');
  let minutes = moment().format('mm');
  let seconds = moment().format('ss');
  setInterval(() => {
    hours = moment().format('HH');
    minutes = moment().format('mm');
    seconds = moment().format('ss');
  }, 1000);
</script>

<div>
  <span style="color: {hoursColor}; font-size: {hoursSize};">{hours}</span>
  <span style="color: {minutesColor}; font-size: {minutesSize};">{minutes}</span
  >
  <span style="color: {secondsColor}; font-size: {secondsSize};">{seconds}</span
  >
</div>

<style>
  div {
    margin: 20px;
  }
</style>
