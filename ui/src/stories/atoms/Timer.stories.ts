import Timer from '../../components/atoms/Timer.svelte';

export default {
  title: 'Atoms/Timer',
  component: Timer,
  tags: ['autodocs']
};

export const Standard = {
  args: {
    hoursSize: '20px',
    minutesSize: '20px',
    secondsSize: '20px',
    hoursColor: 'black',
    minutesColor: 'black',
    secondsColor: 'black'
  }
};

export const Large = {
  args: {
    hoursSize: '40px',
    minutesSize: '40px',
    secondsSize: '40px',
    hoursColor: 'black',
    minutesColor: 'black',
    secondsColor: 'black'
  }
};

export const Red = {
  args: {
    hoursSize: '20px',
    minutesSize: '20px',
    secondsSize: '20px',
    hoursColor: 'red',
    minutesColor: 'red',
    secondsColor: 'red'
  }
};
