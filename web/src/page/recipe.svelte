<script>
    export let recipe;
</script>

<div class="flex">
    <article>
        <h2>{recipe.name}</h2>
        <div class="img-container"><img src={recipe.image} alt /></div>

        <p>{recipe.description}</p>

        Ingredients:
        <ul>
            {#each recipe.ingredients as ingredient}<li>{ingredient}</li>{/each}
        </ul>

        Instructions:
        <ol>
            {#each recipe.instructions as instruction}<li>{instruction}</li>{/each}
        </ol>
    </article>
</div>

<style>
    .flex {
        display: flex;
        justify-content: center;
    }

    article {
        width: 1000px;
    }

    .img-container {
        display: flex;
        justify-content: center;
        padding: 30px;
    }

    img {
        height: 500px;
    }

    h2 {
        text-align: center;
    }
</style>
