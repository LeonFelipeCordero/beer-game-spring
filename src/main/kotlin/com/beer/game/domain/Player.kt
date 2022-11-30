package com.beer.game.domain

import com.beer.game.adapters.out.mongo.PlayerDocument
import com.beer.game.common.Role
import com.beer.game.events.DocumentType
import com.beer.game.events.Event
import com.beer.game.events.EventType
import com.beer.game.events.InternalEventListener
import java.util.UUID

data class Player(
    val id: String = UUID.randomUUID().toString(),
    val name: String,
    val role: Role,
    var stock: Int,
    var backlog: Int,
    var weeklyOrder: Int,
    var lastOrder: Int,
    val cpu: Boolean,
    val orders: MutableList<Order>
) {
    fun addOrder(order: Order) {
        orders.add(order)
    }

    fun emitCreation(listener: InternalEventListener, board: Board) {
        listener.publish(
            Event(
                document = board,
                documentId = board.id,
                entityId = id,
                documentType = DocumentType.PLAYER,
                eventType = EventType.NEW
            )
        )
    }

    fun emitUpdate(listener: InternalEventListener, board: Board) {
        listener.publish(
            Event(
                document = board,
                documentId = board.id,
                entityId = id,
                documentType = DocumentType.PLAYER,
                eventType = EventType.UPDATE
            )
        )
    }
}
