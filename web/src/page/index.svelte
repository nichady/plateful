<script>
    // TODO restrictions such as vegan/vegetarian, allergies, etc...

    import Svelecte from "svelecte";
    import {
        continents as continentMap,
        countries as countryMap,
        getEmojiFlag,
    } from "countries-list";

    const times = ["breakfast", "lunch", "dinner"];
    const flavors = ["sweet", "savory", "salty", "spicy", "sour", "bitter"];
    const continents = Object.values(continentMap).sort();
    const countries = Object.entries(countryMap)
        .sort(([, a], [, b]) => a.name.localeCompare(b.name))
        .map(([id, country]) => ({
            id: country.name,
            name: `${getEmojiFlag(id)}  ${country.name}`,
        }));

    let selectByCountry = false;


    let selectedContinent;
    let selectedCountry;
</script>

<svelte:head>
    <title>Plateful</title>
</svelte:head>

<div class="flex">
    <fieldset>
        <legend><h5>What kinds of meals are you craving?</h5></legend>
        {#each times as t}
            <label>
                <input type="checkbox" />
                {t}
            </label>
        {/each}
    </fieldset>
</div>

<div class="flex">
    <fieldset>
        <legend><h5>What flavors do you want in your meal?</h5></legend>
        {#each flavors as f}
            <label>
                <input type="checkbox" />
                {f}
            </label>
        {/each}
    </fieldset>
</div>

<div class="flex">
    <h5>What ingredients do you want to use?</h5>
    <textarea placeholder="Type ingredients here"></textarea>
</div>

<div class="flex">
    <h5>What ingredients do you want to avoid?</h5>
    <textarea placeholder="Type ingredients here"></textarea>
</div>

<div class="flex">
    <h5>Which types of cuisine do you want?</h5>
    {#if !selectByCountry}
        <Svelecte
            placeholder="Select continents"
            options={continents}
            bind:value={selectedContinent}
            multiple
            style="width: 500px"
        ></Svelecte>
    {:else}
        <Svelecte
            placeholder="Select countries"
            options={countries}
            bind:value={selectedCountry}
            multiple
            style="width: 500px"
        ></Svelecte>
    {/if}
    <label>
        <input type="checkbox" role="switch" bind:checked={selectByCountry} />
        Select by country instead
    </label>
</div>

<div class="flex">
    <button>Generate Recipe</button>
</div>

<style>
    .flex {
        padding: 20px 0;
        display: flex;
        align-items: center;
        flex-direction: column;
    }

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

    textarea {
        width: 500px;
        font-size: 16px;
        resize: none;
    }
</style>
