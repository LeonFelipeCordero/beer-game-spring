package com.beer.game.api

import com.beer.game.adapters.`in`.api.*
import org.springframework.graphql.data.method.annotation.Argument
import org.springframework.graphql.data.method.annotation.MutationMapping
import org.springframework.graphql.data.method.annotation.SchemaMapping
import org.springframework.stereotype.Controller
import reactor.core.publisher.Mono

@Controller
class OrderController(
    private val orderApiAdapter: OrderApiAdapter,
    private val playerApiAdapter: PlayerApiAdapter,
    private val boardApiAdapter: BoardApiAdapter
) {

    @MutationMapping
    fun createOrder(
        @Argument boardId: String,
        @Argument senderId: String,
        @Argument receiverId: String
    ): Mono<OrderGraph> {
        return orderApiAdapter.createOrder(boardId, senderId, receiverId)
    }

    @MutationMapping
    fun deliverOrder(
        @Argument orderId: String,
        @Argument boardId: String,
        @Argument amount: Int? = null
    ): Mono<Response> {
        return orderApiAdapter.deliverOrder(orderId, boardId, amount)
    }

    @SchemaMapping(typeName = "Order", field = "sender")
    fun sender(order: OrderGraph): Mono<PlayerGraph> {
        return playerApiAdapter.getPlayer(order.boardId!!, order.senderId!!)
    }

    @SchemaMapping(typeName = "Order", field = "receiver")
    fun receiver(order: OrderGraph): Mono<PlayerGraph> {
        return playerApiAdapter.getPlayer(order.boardId!!, order.receiverId!!)
    }

    @SchemaMapping(typeName = "Order", field = "board")
    fun board(order: OrderGraph): Mono<BoardGraph> {
        return boardApiAdapter.getBoard(order.boardId!!)
    }
}

