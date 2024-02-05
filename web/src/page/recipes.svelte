<script>
    export let recipes;
</script>

{#if Object.values(recipes).length < 1}
    No recipes yet.
{/if}

<div class="flex">
    {#each Object.entries(recipes) as [u, r]}
        <!-- <article>
            <img src={r.image}>
            <footer><a href={"/recipes/" + u}>{r.name}</a></footer>
        </article> -->
    {/each}
</div>

<style>
    .flex {
        display: flex;
        flex-wrap: wrap;
    }
</style>
