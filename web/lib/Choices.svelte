<script>
    export let choices;
    export let selected = [];
    
    let choiceMap = {};
    for (const c of choices) choiceMap[c] = false;

    $: selected = choices.filter(c => choiceMap[c]);
</script>

<fieldset>
    <legend><slot/></legend>
    {#each choices as choice}
        <label>
            <input type="checkbox" bind:checked={choiceMap[choice]} />
            {choice}
        </label>
    {/each}
</fieldset>

<style>
    fieldset {
        display: flex;
        flex-direction: row;
        gap: 20px;
    }

    legend {
        margin: 0 auto;
    }

    label {
        text-transform: capitalize;
    }
</style>