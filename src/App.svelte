<script lang="ts">
  import { onMount } from 'svelte';
  import { todoService, uuid } from './usecase';

  const ENTER_KEY = 13;
  const ESCAPE_KEY = 27;

  let currentFilter = 'all';
  let editing = null;
  let items = todoService.store();

  onMount(async () => {
    await todoService.list();
  });

  async function updateView() {
    currentFilter = 'all';
    if (window.location.hash === '#/active') {
      currentFilter = 'active';
    } else if (window.location.hash === '#/completed') {
      currentFilter = 'completed';
    }
  }

  window.addEventListener('hashchange', updateView);
  updateView();

  async function clearCompleted() {
    const completedItems = $items.filter((item) => item.completed);
    const promises: Promise<void>[] = [];
    for (const item of completedItems) {
      promises.push(todoService.delete(item));
    }
    await Promise.all(promises);
  }

  async function remove(index: number) {
    const item = $items[index];
    await todoService.delete(item);
  }

  async function toggleAll(event: Event) {
    const target = event.target as HTMLInputElement;
    const newItems = $items.map((item) => ({
      ...item,
      completed: target.checked,
    }));
    const promises: Promise<void>[] = [];
    for (const item of newItems) {
      promises.push(todoService.update(item));
    }
    await Promise.all(promises);
  }

  async function toggle(index: number) {
    const item = $items[index];
    await todoService.update(item);
  }

  async function createNew(event: KeyboardEvent) {
    const target = event.target as HTMLInputElement;
    if (event.which === ENTER_KEY) {
      await todoService.create({
        id: uuid(),
        description: target.value,
        completed: false,
      });
      target.value = '';
    }
  }

  function handleEdit(event: KeyboardEvent) {
    const target = event.target as HTMLInputElement;
    if (event.which === ENTER_KEY) target.blur();
    else if (event.which === ESCAPE_KEY) editing = null;
  }

  async function submit(event: FocusEvent) {
    const target = event.target as HTMLInputElement;
    const item = $items[editing];
    item.description = target.value;
    editing = null;
    await todoService.update(item);
  }

  $: filtered =
    currentFilter === 'all'
      ? $items
      : currentFilter === 'completed'
      ? $items.filter((item) => item.completed)
      : $items.filter((item) => !item.completed);

  $: numActive = $items.filter((item) => !item.completed).length;

  $: numCompleted = $items.filter((item) => item.completed).length;
</script>

<header class="header">
  <h1>todos</h1>
  <!-- svelte-ignore a11y-autofocus -->
  <input class="new-todo" on:keydown={createNew} placeholder="What needs to be done?" autofocus />
</header>

{#if $items.length > 0}
  <section class="main">
    <input
      id="toggle-all"
      class="toggle-all"
      type="checkbox"
      on:change={toggleAll}
      checked={numCompleted === $items.length} />
    <label for="toggle-all">Mark all as complete</label>

    <ul class="todo-list">
      {#each filtered as item, index (item.id)}
        <li class="{item.completed ? 'completed' : ''} {editing === index ? 'editing' : ''}">
          <div class="view">
            <input class="toggle" type="checkbox" bind:checked={item.completed} on:change={() => toggle(index)} />
            <label on:dblclick={() => (editing = index)}>{item.description}</label>
            <button on:click={() => remove(index)} class="destroy" />
          </div>

          <!-- svelte-ignore a11y-autofocus -->
          {#if editing === index}
            <input value={item.description} id="edit" class="edit" on:keydown={handleEdit} on:blur={submit} autofocus />
          {/if}
        </li>
      {/each}
    </ul>

    <footer class="footer">
      <span class="todo-count">
        <strong>{numActive}</strong>
        {numActive === 1 ? 'item' : 'items'} left
      </span>

      <ul class="filters">
        <li>
          <a class={currentFilter === 'all' ? 'selected' : ''} href="#/">All</a>
        </li>
        <li>
          <a class={currentFilter === 'active' ? 'selected' : ''} href="#/active">Active</a>
        </li>
        <li>
          <a class={currentFilter === 'completed' ? 'selected' : ''} href="#/completed">Completed</a>
        </li>
      </ul>

      {#if numCompleted}
        <button class="clear-completed" on:click={clearCompleted}>Clear completed</button>
      {/if}
    </footer>
  </section>
{/if}
