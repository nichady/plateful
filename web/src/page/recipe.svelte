<script>
    export let name;
    export let image;
    export let body;
</script>

<div class="flex">
    <article>
        <h2>{name}</h2>
        <div class="img-container">
            <img src={image} alt />
        </div>
        <p>{body}</p>
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

    p {
        white-space: pre-wrap;
        font-size: 16px;
    }
</style>
