package com.beer.game.domain

import com.beer.game.common.BoardState
import com.beer.game.events.DocumentType
import com.beer.game.events.Event
import com.beer.game.events.EventType
import com.beer.game.events.InternalEventListener
import java.time.LocalDateTime

class Board(
    val id: String,
    val name: String,
    var state: BoardState,
    var full: Boolean,
    var finished: Boolean,
    val createdAt: LocalDateTime,
    val players: MutableList<Player>,
) {

    fun addPlayer(player: Player) {
        players.add(player)
    }

    fun findPlayer(playerId: String): Player {
        return players
            .first { it.id == playerId }
    }

    fun findOrder(orderId: String): Order {
        return players
            .flatMap { it.orders }
            .first { it.id == orderId }
    }

    fun emitUpdate(listener: InternalEventListener) {
        listener.publish(
            Event(
                document = this,
                documentId = id,
                documentType = DocumentType.PLAYER,
                eventType = EventType.UPDATE
            )
        )
    }
}
