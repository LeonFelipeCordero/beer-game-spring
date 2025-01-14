package com.beer.game.api.order

import com.beer.game.api.domain.Response
import com.beer.game.application.board.BoardService
import com.beer.game.application.order.OrderEvenListener
import com.beer.game.application.order.OrderService
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import reactor.core.scheduler.Schedulers

@Component
class OrderApiAdapter(
    private val boardService: BoardService,
    private val orderService: OrderService,
    private val orderEvenListener: OrderEvenListener,
) {

    fun createOrder(
        receiverId: String,
    ): Mono<OrderGraph> {
        return Mono.fromCallable {
            orderService.createOrder(receiverId)
        }.map {
            OrderGraph.fromOrder(it.first, it.second, receiverId)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun deliverOrder(
        orderId: String,
        amount: Int? = null,
    ): Mono<Response> {
        return Mono.fromCallable {
            orderService.deliverOrder(orderId, amount)
        }.map {
            Response(message = "Order delivered", status = 200)
        }
    }

    fun newOrderSubscription(playerId: String): Flux<OrderGraph> {
        val board = boardService.loadBoardByPlayerId(playerId)
        return orderEvenListener.subscribeNewOrder(playerId)
            .map { OrderGraph.fromOrder(it, board.id, it.receiver) }
    }

    fun orderDeliverySubscription(playerId: String): Flux<OrderGraph> {
        val board = boardService.loadBoardByPlayerId(playerId)
        return orderEvenListener.subscribeUpdateDelivery(playerId)
            .map { OrderGraph.fromOrder(it, board.id, it.receiver) }
    }
}
