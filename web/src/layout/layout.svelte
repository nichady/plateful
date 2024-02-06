<script>
    import '@fontsource/open-sans';
    import "@picocss/pico";
    import "../../app.css";

    import plate from "../../lib/plate.svg";
</script>

<header class="container-fluid">
    <img src={plate} alt="Icon of plate" />
    <h2 class="container-fluid">Plateful - The AI-Powered Recipe Generator.</h2>
    <img src={plate} alt="Icon of plate" />
</header>

<main>
    <slot />
</main>

<style>
    header {
        display: flex;
        align-items: center;
        padding: 10px 20px;
        height: 110px;
        background: var(--gold);
        box-shadow: rgba(0, 0, 0, 0.25) 0px 2px 2px;
    }
    
    img {
        object-fit: contain;
        height: 100%;
    }
    
    h2 {
        text-align: center;
        flex: 1;
    }

    main {
        padding: 20px;
    }
</style>
