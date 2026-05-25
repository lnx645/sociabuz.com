<script lang="ts">
  import { mq } from "@/core/utils";
  import DuolingoIcon from "@/icons/duolingo-icon.svelte";
  import IconHome from "@/icons/icon-home.svelte";
  import { css } from "@emotion/css";
  import NavItem from "./nav-item.svelte";
  import MoreIcon from "@/icons/more-icon.svelte";
  import ListIcon from "@/icons/list-icon.svelte";
  import FeaturesIcon from "@/icons/features-icon.svelte";
  const sidebar = {
    menu: css({
      display: "flex",
      flexDirection: "column",
      gap: 10,
      padding: 10,
    }),
    wrapper: css(
      mq({
        maxWidth: 260,
        minWidth: 260,
        borderRight: "2px solid",
        borderColor: "#e5e5e5",
        height: "100vh",
        top: ["56px", 0],
        display: ["block"],
        position: ["fixed", "relative"],
        background: "white",
        zIndex: 99,
      }),
    ),
    brand: css(
      mq({
        display: ["none", "inline-block"],
        paddingTop: 24,
        paddingInline: 10,
        paddingBottom: 5,
        fontWeight: "bold",
        svg: {
          width: "50%",
        },
      }),
    ),
  };
</script>

<aside class={sidebar.wrapper}>
  <div class={sidebar.brand}>
    <DuolingoIcon />
  </div>
  <div class={sidebar.menu}>
    <NavItem
      isActive={true}
      title="PROFILE"
      description="@dadanhdyt - creator"
      icon={IconHome}
    />
    <NavItem title="TRANSAKSI" description="Kelola transaksi" icon={ListIcon} />
    <NavItem
      title="FEATURES"
      description="Overlay,SondBoard,dll"
      icon={FeaturesIcon}
    />
    <NavItem title="MORE" icon={MoreIcon} />
  </div>
</aside>
