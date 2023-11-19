export function clickOutside(
  node: HTMLElement,
  handler: (arg: HTMLElement) => void,
) {
  const handleClick: Parameters<
    typeof document.addEventListener<"click">
  >[1] = (event) => {
    if (
      event.target &&
      node &&
      !node.contains(event.target as Node) &&
      !event.defaultPrevented
    ) {
      handler(node);
    }
  };

  document.addEventListener("click", handleClick, true);

  return {
    destroy() {
      document.removeEventListener("click", handleClick, true);
    },
  };
}
