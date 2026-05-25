<script lang="ts">
  import { flexbox } from "@/styled";
  import { css, cx } from "@emotion/css";
  import { Avatar } from "bits-ui";
</script>

<Avatar.Root
  class={cx(
    css({
      width: 90,
      height: 90,
      border: "2px solid var(--color-disabled)",
      borderRadius: 100,
    }),
    flexbox.flex,
    flexbox.alignItemsCenter,
    flexbox.justifyCenter,
  )}
>
  <Avatar.Fallback>D</Avatar.Fallback>
</Avatar.Root>
