import Button from '../../components/atoms/Button.svelte';

export default {
  title: 'Atoms/Button',
  component: Button,
  tags: ['autodocs']
};

export const Standard = {
  args: {
    name: 'Button',
    width: '100px',
    height: '30px',
    radius: '10px',
    fetchData: () => {
      console.log('log');
    }
  }
};

export const Medium = {
  args: {
    name: 'Button',
    width: '120px',
    height: '40px',
    radius: '10px',
    fetchData: () => {
      console.log('log');
    }
  }
};

export const Large = {
  args: {
    name: 'Button',
    width: '200px',
    height: '60px',
    radius: '10px',
    fetchData: () => {
      console.log('log');
    }
  }
};
