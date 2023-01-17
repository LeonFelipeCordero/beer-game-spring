import {Board, Player} from "../../gql/graphql";
import {Show} from "solid-js";

function GameStatus(props: { player: Player, board: Board, createOrder: Function }) {
    return (
        <div class="mt-5">
            <strong>Current status</strong>
            <div class="overflow-x-auto">
                <table class="table w-full">
                    <thead>
                    <tr>
                        <th>Stock</th>
                        <th>Last week</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <Show when={props.player.stock > 10} keyed
                              fallback={
                                  <td>
                                      {props.player.stock}
                                  </td>
                              }>
                            <td>
                                {props.player.stock}
                            </td>
                        </Show>
                        <td>
                            {props.player.lastOrder}
                        </td>
                    </tr>
                    </tbody>
                </table>
            </div>
            <form phx-submit="create_order" class="flex flex-col mt-5">
                <Show when={props.player.role == "FACTORY"} keyed
                      fallback={
                          <label for="amount" class="text-gray-700 text-m font-bold mb-2">
                              Weekly order
                          </label>
                      }
                >
                    <label for="amount" class="text-gray-700 text-m font-bold mb-2">
                        Weekly production
                    </label>
                </Show>
                <input class="border rounded p-2 text-gray-700"
                       type="number"
                       name="amount"
                       value={props.player.weeklyOrder}/>
                <Show when={props.player.role != "FACTORY"} keyed>
                    <input
                        class="bg-blue-500 hover:bg-blue-700 text-white font-bold p-2 rounded mt-2"
                        type="submit"
                        value="order"
                        onClick={(e) => {
                            e.preventDefault()
                            props.createOrder(props.board.id, props.player.id)
                        }}
                    />
                </Show>
            </form>
        </div>
    )
}

export default GameStatus
