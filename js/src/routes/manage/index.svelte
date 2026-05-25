<script lang="ts">
  import LinkIcon from "@/icons/link-icon.svelte";
  import { css, cx } from "@emotion/css";
  import { card, flexbox, fontSize, typography } from "@/styled";
  import colors from "@/styled/colors";
</script>

<img
  src="https://dy8n3onijof8f.cloudfront.net/static/images/wp/03740b1e7e77ad332349.svg"
  alt=""
/>

<div
  class={css({
    padding: 10,
  })}
>
  <div class={card.wrapper}>
    <div class={card.content}>
      <div
        class={cx(
          flexbox.flex,
          typography.fontBold,
          flexbox.alignItemsCenter,
          flexbox.gapXS,
        )}
      >
        <LinkIcon />
        <span>Link Halaman Kamu</span>
      </div>
      <a
        class={cx(
          typography.underline,
          typography.fontSemiBold,
          colors.primary,
          fontSize.textXs,
        )}
        href="https://sobuzz.id/dadanhdyt">https://sobuzz.id/dadanhdyt</a
      >
    </div>
  </div>
</div>
