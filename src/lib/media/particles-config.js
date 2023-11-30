const particlesConfig = {
  background: {
    color: {
      value: "#120f0f",
    },
  },
  fpsLimit: 120,
  interactivity: {

    modes: {
      push: {
        quantity: 0.5,
      },
      repulse: {
        distance: 400,
        duration: 0.1,
      },
    },
  },
  particles: {
    color: {
      value: ["rgb(206, 46, 238)", "rgb(15, 195, 233)"],
    },
    links: {
      color: ["rgb(206, 46, 238)", "rgb(15, 195, 233)", "rgb(206, 46, 238)"],
      distance: 100,
      enable: true,
      opacity: 0.5,
      width: 2,
    },
    move: {
      direction: "none",
      enable: true,
      outModes: {
        default: "bounce",
      },
      random: false,
      speed: 1,
      straight: false,
    },
    number: {
      density: {
        enable: true,
        area: 600,
      },
      value: 60,
    },
    opacity: {
      value: 0.5,
    },
    shape: {
      type: "circle",
    },
    size: {
      value: { min: 1, max: 2 },
    },
  },
  detectRetina: true,
}

export default particlesConfig