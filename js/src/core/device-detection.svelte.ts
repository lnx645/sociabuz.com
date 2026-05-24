export let screenSize = $state({
  isMobile: false,
  isTablet: false,
  isDekstop: false,
});
let breakpoints = {
  sm: 576,
  md: 768,
  lg: 992,
  xl: 1200,
};
function mobileDetection() {
  const width = window.innerWidth;

  screenSize.isMobile = width <= breakpoints.sm;
  screenSize.isTablet = width > breakpoints.sm && width <= breakpoints.md;
  screenSize.isDekstop = width > breakpoints.md;
}

export function InitializeScreenSize() {
  $effect(() => {
    mobileDetection();
    window.addEventListener("resize", mobileDetection);
    return () => {
      window.removeEventListener("resize", mobileDetection);
    };
  });
}

export default screenSize;
